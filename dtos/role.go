package dtos

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
