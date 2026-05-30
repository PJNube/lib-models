package dtos

type CreateBackupRequest struct {
	Name                 string `json:"name"`
	Tier                 string `json:"tier"`
	IncludeSystemUpgrade bool   `json:"includeSystemUpgrade"`
}

type DeleteBackupsRequest struct {
	UUIDs []string `json:"uuids"`
}

type ExtensionVersionLock struct {
	Id      string `json:"id"`
	Version string `json:"version"`
}
