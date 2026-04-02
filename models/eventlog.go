package models

import (
	"time"

	"github.com/PJNube/lib-models/utils/pjnjson"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type EventLog struct {
	ID          uint64                      `json:"id" gorm:"primaryKey;autoIncrement"`
	ExtensionID string                      `json:"extensionId" gorm:"index;not null"`
	Tags        datatypes.JSONSlice[string] `json:"tags" gorm:"type:text"`
	Payload     datatypes.JSON              `json:"payload" gorm:"type:text"`
	CreatedAt   time.Time                   `json:"createdAt,omitempty" gorm:"not null;autoCreateTime"`
}

func (f *EventLog) BeforeCreate(tx *gorm.DB) (err error) {
	f.Payload = pjnjson.MarshalJson(f.Payload)
	return
}

func (f *EventLog) BeforeUpdate(tx *gorm.DB) (err error) {
	f.Payload = pjnjson.MarshalJson(f.Payload)
	return
}
