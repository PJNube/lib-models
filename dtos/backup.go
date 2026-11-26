package dtos

type CreateBackupRequest struct {
	Name string `json:"name"`
	Tier string `json:"tier"`
	// Components is required for the custom tier and refused for every other.
	Components []string `json:"components,omitempty"`
}

type DeleteBackupsRequest struct {
	UUIDs []string `json:"uuids"`
}

type ExtensionVersionLock struct {
	Version     string `json:"version"`
	VersionPath string `json:"version_path"`
	DataPath    string `json:"data_path"`
}
