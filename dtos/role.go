package dtos

type Role struct { // role create and upsert
	Name        string      `json:"name"`
	ExtensionID string      `json:"extensionId,omitempty"`
	Resources   []*Resource `json:"resources,omitempty"`
}

type RoleResponse struct { // hierarchical role response : role -> extensions -> resources
	Name       string       `json:"name"`
	Extensions []*Extension `json:"extensions,omitempty"`
}

type Extension struct {
	ID        string      `json:"id"`
	Resources []*Resource `json:"resources,omitempty"`
}

type Resource struct {
	ID     string `json:"id"`
	Action string `json:"action,omitempty"`
}

type CasbinPolicy struct {
	Role   string `json:"role"`
	Source string `json:"source"` // combination of extensionID and resourceID
	Topics string `json:"topics"`
	Action string `json:"action"`
}
