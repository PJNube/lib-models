package dtos

type Permission struct {
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	ExtensionId string `json:"extensionId"`
}

type Permissions struct {
	Resources []*Permission `json:"resources"`
}

type FeatureResource struct {
	ID        string   `json:"id"`
	Resources []string `json:"resources,omitempty"`
}

type ResourceParam struct {
	Ids         []string `json:"ids"`
	ExtensionId string   `json:"extensionId"`
}
