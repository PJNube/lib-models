package models

import (
	"time"

	"github.com/PJNube/lib-models/datatypes"
	"github.com/PJNube/lib-models/utils/nuuid"
	"gorm.io/gorm"
)

type Backup struct {
	UUID      string                   `json:"uuid" sql:"uuid" gorm:"type:varchar(255);unique;primaryKey"`
	Name      string                   `json:"name" gorm:"unique;not null"`
	FileName  string                   `json:"fileName"`
	Size      int64                    `json:"size"`
	Status    datatypes.ProgressStatus `json:"status"`
	Progress  int                      `json:"progress"`
	Message   string                   `json:"message,omitempty"`
	CreatedAt time.Time                `json:"createdAt"`
	Tier      string                   `json:"tier,omitempty"`
}

func (b *Backup) BeforeCreate(tx *gorm.DB) (err error) {
	if b.UUID == "" {
		b.UUID = nuuid.ShortUUID("bac")
	}
	return
}
