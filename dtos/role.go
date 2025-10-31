package dtos

type Role struct {
	Name   string   `json:"name"`
	Source string   `json:"source"`
	Topics []string `json:"topics"`
	Effect *string  `json:"effect"`
}

type CasbinPolicy struct {
	Role   string `json:"role"`
	Source string `json:"source"` // combination of extensionID and resourceID
	Topics string `json:"topics"`
	Effect string `json:"effect"`
}

type RoleRequest struct {
	Name        string   `json:"name"`
	ExtensionID string   `json:"extensionId"`
	ResourceIds []string `json:"resourceIds"`
	Effect      *string  `json:"effect"`
}
