package models

import (
	"github.com/PJNube/lib-models/utils/nuuid"
	"gorm.io/gorm"
)

// Folder represents a hierarchical folder structure
type Folder struct {
	UUID       string   `json:"uuid" gorm:"type:varchar(255);unique;primaryKey"`
	Name       string   `json:"name" gorm:"type:varchar(255);not null;uniqueIndex:idx_name_parent_uuid"`
	ParentUUID *string  `json:"parentUuid,omitempty" gorm:"type:varchar(255);default:null;uniqueIndex:idx_name_parent_uuid"`
	Folders    []Folder `json:"folders,omitempty" gorm:"foreignKey:ParentUUID;constraint:OnDelete:CASCADE"`
	Hosts      []Host   `json:"hosts,omitempty" gorm:"foreignKey:FolderUUID;constraint:OnDelete:CASCADE"`
}

func (f *Folder) BeforeCreate(tx *gorm.DB) (err error) {
	if f.UUID == "" {
		f.UUID = nuuid.ShortUUID("fol")
	}
	return
}
