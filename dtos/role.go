package dtos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"slices"
)

type Role struct { // role create and upsert
	Name         string      `json:"name"`
	Action       string      `json:"action,omitempty"`
	IsPredefined bool        `json:"isPredefined"`
	Resources    []*Resource `json:"resources,omitempty"`
	CE           CEGrants    `json:"ce"`

	// Categories lists a role's CE grants one category per row, for the Roles
	// grid. On a read, CategoriesUnavailable says CE didn't answer, so only the
	// categories the role already grants are listed.
	Categories            []CategoryGrant `json:"categories,omitzero"`
	CategoriesUnavailable bool            `json:"categoriesUnavailable,omitempty"`
}

// CEGrants lists, per verb and level, the Control Engine categories a role
// reaches. Category names only: CE assigns the bits per engine.
type CEGrants struct {
	Read   CEGrantLevels `json:"read"`
	Write  CEGrantLevels `json:"write"`
	Invoke CEGrantLevels `json:"invoke"`
}

type CEGrantLevels struct {
	Admin    []string `json:"admin"`
	Operator []string `json:"operator"`
}

// Lists returns the six category lists, one per verb and level, so callers
// can read or replace each one without naming them all.
func (g *CEGrants) Lists() []*[]string {
	return []*[]string{
		&g.Read.Admin, &g.Read.Operator,
		&g.Write.Admin, &g.Write.Operator,
		&g.Invoke.Admin, &g.Invoke.Operator,
	}
}

// UnmarshalJSON refuses unknown keys at any depth: a misspelt verb or level
// would otherwise be dropped silently and save a narrower grant than was sent.
func (g *CEGrants) UnmarshalJSON(data []byte) error {
	type plain CEGrants
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	var p plain
	if err := dec.Decode(&p); err != nil {
		return err
	}
	*g = CEGrants(p)
	return nil
}

const (
	CELevelNone     = "none"
	CELevelOperator = "operator"
	CELevelAdmin    = "admin"
)

// CategoryGrant is the level a role holds on one category for each verb. An
// admin grant also covers operator-level parts, so one level per verb is enough.
type CategoryGrant struct {
	Name   string `json:"name"`
	Read   string `json:"read"`
	Write  string `json:"write"`
	Invoke string `json:"invoke"`
}

// UnmarshalJSON refuses unknown keys and any verb whose level is missing or
// unknown: saving those as "none" would silently remove access.
func (c *CategoryGrant) UnmarshalJSON(data []byte) error {
	type plain CategoryGrant
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	var p plain
	if err := dec.Decode(&p); err != nil {
		return err
	}
	grant := CategoryGrant(p)
	if err := grant.validateLevels(); err != nil {
		return err
	}
	*c = grant
	return nil
}

func (c CategoryGrant) validateLevels() error {
	for _, level := range []string{c.Read, c.Write, c.Invoke} {
		switch level {
		case CELevelNone, CELevelOperator, CELevelAdmin:
		default:
			return fmt.Errorf("category %q: level must be none, operator or admin, got %q", c.Name, level)
		}
	}
	return nil
}

// Category returns the levels held on one category, as a CategoryGrant.
func (g CEGrants) Category(name string) CategoryGrant {
	return CategoryGrant{
		Name:   name,
		Read:   g.Read.levelOf(name),
		Write:  g.Write.levelOf(name),
		Invoke: g.Invoke.levelOf(name),
	}
}

// SetCategory replaces the levels held on one category and keeps the others.
func (g *CEGrants) SetCategory(c CategoryGrant) error {
	if err := c.validateLevels(); err != nil {
		return err
	}
	for _, list := range g.Lists() {
		*list = slices.DeleteFunc(*list, func(name string) bool { return name == c.Name })
	}
	g.Read.add(c.Read, c.Name)
	g.Write.add(c.Write, c.Name)
	g.Invoke.add(c.Invoke, c.Name)
	return nil
}

func (l CEGrantLevels) levelOf(name string) string {
	if slices.Contains(l.Admin, name) {
		return CELevelAdmin
	}
	if slices.Contains(l.Operator, name) {
		return CELevelOperator
	}
	return CELevelNone
}

func (l *CEGrantLevels) add(level, name string) {
	switch level {
	case CELevelAdmin:
		l.Admin = append(l.Admin, name)
	case CELevelOperator:
		l.Operator = append(l.Operator, name)
	}
}

type Resource struct {
	ExtensionID string `json:"extensionId,omitempty"`
	ID          string `json:"id"`
	Action      string `json:"action,omitempty"`
}

type CasbinPolicy struct {
	Name   string `json:"name"`
	Source string `json:"source"`
	Topics string `json:"topics"`
	Action string `json:"action"`
}
