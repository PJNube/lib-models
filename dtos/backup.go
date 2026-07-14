package dtos

type CreateBackupRequest struct {
	Name string `json:"name"`
	Tier string `json:"tier"`
}

type DeleteBackupsRequest struct {
	UUIDs []string `json:"uuids"`
}

type ExtensionVersionLock struct {
	Id      string `json:"id"`
	Version string `json:"version"`
}
