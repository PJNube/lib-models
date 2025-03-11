package models

import (
	"github.com/pjnube/lib-models/utils/nuuid"
	"gorm.io/gorm"
)

type Host struct {
	UUID       string  `json:"uuid" gorm:"type:varchar(255);unique;primaryKey"`
	Name       string  `json:"name" gorm:"type:varchar(255);not null;uniqueIndex:idx_name_folder_uuid"`
	GlobalUUID string  `json:"globalUuid" gorm:"type:varchar(255);not null"`
	FolderUUID *string `json:"folderUuid" gorm:"type:varchar(255);default:null;uniqueIndex:idx_name_folder_uuid,expression:IFNULL(folder_uuid\\,\"\")"`
}

func (f *Host) BeforeCreate(tx *gorm.DB) (err error) {
	if f.UUID == "" {
		f.UUID = nuuid.ShortUUID("hos")
	}
	return
}
