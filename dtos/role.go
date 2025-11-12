package dtos

type Role struct { // role create and upsert
	Name      string      `json:"name"`
	Action    string      `json:"action,omitempty"`
	Resources []*Resource `json:"resources,omitempty"`
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
