package models

import (
	"time"

	"github.com/PJNube/lib-models/datatypes"
	"github.com/PJNube/lib-models/utils/nuuid"
	"gorm.io/gorm"
)

type BackupExtension struct {
	Extension bool `json:"extension"`
	Data      bool `json:"data"`
}
type Backup struct {
	UUID       string                     `json:"uuid" sql:"uuid" gorm:"type:varchar(255);unique;primaryKey"`
	Name       string                     `json:"name" gorm:"unique;not null"`
	FileName   string                     `json:"fileName"`
	Extensions map[string]BackupExtension `json:"extensions,omitempty" gorm:"type:json;serializer:json"`
	Size       int64                      `json:"size"`
	Status     datatypes.ProgressStatus   `json:"status"`
	Message    string                     `json:"message,omitempty"`
	CreatedAt  time.Time                  `json:"createdAt"`
}

func (b *Backup) BeforeCreate(tx *gorm.DB) (err error) {
	if b.UUID == "" {
		b.UUID = nuuid.ShortUUID("bac")
	}
	return
}
