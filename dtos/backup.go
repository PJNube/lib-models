package dtos

import "github.com/PJNube/lib-models/models"

type CreateBackupRequest struct {
	Name                 string                            `json:"name"`
	BackupExtensions     map[string]models.BackupExtension `json:"backupExtensions"`
	IncludeSystemUpgrade bool                              `json:"includeSystemUpgrade"`
}

type DeleteBackupsRequest struct {
	UUIDs []string `json:"uuids"`
}

type ExtensionVersionLock struct {
	EnvelopeId  string `json:"envelopeId"`
	Version     string `json:"version"`
	VersionPath string `json:"version_path"`
}
