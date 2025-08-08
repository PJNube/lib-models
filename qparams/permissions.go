package qparams

type FeatureParam struct {
	ExtensionId string    `json:"extensionId"`
	Username    string    `json:"username"`
	Features    []Feature `json:"features"`
}

type Feature struct {
	Feature string   `json:"feature"`
	Action  []string `json:"action"`
}
