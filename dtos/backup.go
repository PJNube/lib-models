package dtos

import (
	"fmt"

	"github.com/PJNube/lib-models/datatypes"
)

type CreateBackupRequest struct {
	Name       string   `json:"name"`
	Components []string `json:"components"`
}

func (c *CreateBackupRequest) ValidateComponents() (bool, error) {
	for _, component := range c.Components {
		if _, ok := datatypes.BackupComponentMap[component]; !ok {
			return false, fmt.Errorf("invalid component type %s", component)
		}
	}

	return true, nil
}

type DeleteBackupsRequest struct {
	UUIDs []string `json:"uuids"`
}
