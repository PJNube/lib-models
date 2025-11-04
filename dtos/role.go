package dtos

type Role struct {
	Name        string        `json:"name"`
	ExtensionID string        `json:"extensionId,omitempty"`
	ResourceIds []*ResourceID `json:"resourceIds,omitempty"`
}
type CasbinPolicy struct {
	Role   string `json:"role"`
	Source string `json:"source"` // combination of extensionID and resourceID
	Topics string `json:"topics"`
	Action string `json:"action"`
}

type ResourceID struct {
	ID     string `json:"id"`
	Action string `json:"action,omitempty"`
}

type RoleResponse map[string]map[string][]ResourceID
