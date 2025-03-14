package dtos

import "github.com/PJNube/lib-models/models"

type FoldersHosts struct {
	Folders []models.Folder `json:"folders"`
	Hosts   []models.Host   `json:"hosts"`
}
