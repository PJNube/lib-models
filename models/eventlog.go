package models

import (
	"time"

	"github.com/PJNube/lib-models/utils/nuuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type EventLog struct {
	UUID        string                      `json:"uuid" gorm:"type:text;primaryKey;not null"`
	ExtensionID string                      `json:"extensionId" gorm:"index;not null"`
	Timestamp   time.Time                   `json:"timestamp" gorm:"index;not null;autoCreateTime"`
	Tags        datatypes.JSONSlice[string] `json:"tags" gorm:"type:text"`
	Payload     datatypes.JSON              `json:"payload" gorm:"type:text"`
}

func (e *EventLog) BeforeCreate(tx *gorm.DB) (err error) {
	if e.UUID == "" {
		e.UUID = nuuid.ShortUUID("evt")
	}
	return
}
