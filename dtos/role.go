package dtos

import (
	"bytes"
	"encoding/json"
)

type Role struct { // role create and upsert
	Name         string      `json:"name"`
	Action       string      `json:"action,omitempty"`
	IsPredefined bool        `json:"isPredefined"`
	Resources    []*Resource `json:"resources,omitempty"`
	CE           CEGrants    `json:"ce"`
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
