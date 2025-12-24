package dtos

type CreateBackupRequest struct {
	Name       string   `json:"name"`
	Components []string `json:"components"`
}

type DeleteBackupsRequest struct {
	UUIDs []string `json:"uuids"`
}

type ExtensionVersionLock struct {
	Version     string `json:"version"`
	VersionPath string `json:"version_path"`
	DataPath    string `json:"data_path"`
}
