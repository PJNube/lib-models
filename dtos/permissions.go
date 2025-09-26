package dtos

type Permission struct {
	Item          string `json:"item"` // resource
	Action        string `json:"action"`
	ExtensionName string `json:"extensionId"`
}

type Permissions struct {
	Resources []*Permission `json:"items"`
}

type FeatureResource struct {
	ID          string   `json:"id"`
	Description string   `json:"description,omitempty"`
	Items       []string `json:"items,omitempty"` // resources
}

type ResourceParam struct {
	Ids          []string `json:"ids"`
	ExtensionId  string   `json:"extensionId"`  // now full id with profile, vendor and name for casbin
	Version      string   `json:"version"`      // v1.0.0
	Architecture string   `json:"architecture"` // arm64, amd64
}
