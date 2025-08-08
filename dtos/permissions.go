package dtos

type Permission struct {
	Resource string   `json:"resource"`
	Action   []string `json:"action"`
	Feature  string   `json:"feature"`
	Source   string   `json:"source"`
}

type Permissions struct {
	Resources []*Permission `json:"resources"`
}

type ExtensionPermissions struct {
	Resources []*Permission `json:"resources"`
	Features  []string      `json:"features"`
}
