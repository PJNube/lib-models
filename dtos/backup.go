package dtos

type CreateBackupRequest struct {
	Name                 string   `json:"name"`
	Components           []string `json:"components"`
	IncludeSystemUpgrade bool     `json:"includeSystemUpgrade"`
}

type DeleteBackupsRequest struct {
	UUIDs []string `json:"uuids"`
}

type ExtensionVersionLock struct {
	Id      string `json:"id"`
	Version string `json:"version"`
}
