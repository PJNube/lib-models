package models

import (
	"time"

	"github.com/PJNube/lib-models/datatypes"
	"github.com/PJNube/lib-models/utils/nuuid"
	"gorm.io/gorm"
)

type Restore struct {
	UUID                 string                   `json:"uuid" sql:"uuid" gorm:"type:varchar(255);unique;primaryKey"`
	BackupUUID           string                   `json:"backupUuid" gorm:"type:varchar(255)"`
	BackupFileName       string                   `json:"backupFileName,omitempty"`
	Tier                 string                   `json:"tier,omitempty"`
	IncludeSystemUpgrade bool                     `json:"includeSystemUpgrade,omitempty"`
	Status               datatypes.ProgressStatus `json:"status"`
	Progress             int                      `json:"progress"`
	Message              string                   `json:"message,omitempty"`
	CreatedAt            time.Time                `json:"createdAt"`
}

func (b *Restore) BeforeCreate(tx *gorm.DB) (err error) {
	if b.UUID == "" {
		b.UUID = nuuid.ShortUUID("res")
	}
	return
}
