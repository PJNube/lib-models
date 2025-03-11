package dtos

import "github.com/PJNube/lib-models/models"

type FoldersHosts struct {
	Folders []models.Folder `json:"folders,omitempty"`
	Hosts   []models.Host   `json:"hosts,omitempty"`
}
