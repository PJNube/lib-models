package models

import (
	"time"

	"github.com/PJNube/lib-models/datatypes"
	"github.com/PJNube/lib-models/utils/nuuid"
	gormDatatypes "gorm.io/datatypes"
	"gorm.io/gorm"
)

type Backup struct {
	UUID       string                          `json:"uuid" sql:"uuid" gorm:"type:varchar(255);unique;primaryKey"`
	Name       string                          `json:"name" gorm:"not null"`
	FileName   string                          `json:"fileName"`
	Components gormDatatypes.JSONSlice[string] `json:"components,omitempty"`
	Size       string                          `json:"size"`
	Status     datatypes.ProgressStatus        `json:"status"`
	Message    string                          `json:"message,omitempty"`
	CreatedAt  time.Time                       `json:"createdAt"`
}

func (b *Backup) BeforeCreate(tx *gorm.DB) (err error) {
	if b.UUID == "" {
		b.UUID = nuuid.ShortUUID("bac")
	}
	return
}
