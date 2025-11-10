package dtos

type Role struct { // role create and upsert
	Name        string      `json:"name"`
	ExtensionID string      `json:"extensionId,omitempty"`
	Resources   []*Resource `json:"resources,omitempty"`
}

type Resource struct {
	ID     string `json:"id"`
	Action string `json:"action,omitempty"`
}

type CasbinPolicy struct {
	Name   string `json:"name"`
	Source string `json:"source"`
	Topics string `json:"topics"`
	Action string `json:"action"`
}
