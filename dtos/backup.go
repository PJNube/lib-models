package dtos

import (
	"fmt"

	"github.com/PJNube/lib-models/datatypes"
)

type CreateBackupRequest struct {
	Name       string                      `json:"name"`
	Components []datatypes.BackupComponent `json:"components"`
}

func (c *CreateBackupRequest) ValidateComponents() (bool, error) {
	for _, component := range c.Components {
		if _, ok := datatypes.BackupComponentMap[component]; !ok {
			return false, fmt.Errorf("invalid component type %s", component)
		}
	}

	return true, nil
}

func (c *CreateBackupRequest) ComponentsAsStringSlice() []string {
	components := make([]string, 0, len(c.Components))
	for _, component := range c.Components {
		components = append(components, string(component))
	}
	return components
}

type DeleteBackupsRequest struct {
	UUIDs []string `json:"uuids"`
}

type ExtensionVersionLock struct {
	Version     string `json:"version"`
	VersionPath string `json:"version_path"`
	DataPath    string `json:"data_path"`
}
