package models

import (
	"encoding/json"

	"github.com/PJNube/lib-models/utils/nuuid"
	"gorm.io/gorm"
)

type UserLayout struct {
	UUID     string          `gorm:"primaryKey;type:VARCHAR(255)" json:"uuid"`
	Name     string          `gorm:"type:VARCHAR(255);not null;uniqueIndex" json:"name"`
	UserUUID string          `gorm:"index;not null" json:"userUuid,omitempty"`
	Layout   json.RawMessage `gorm:"type:LONGBLOB;not null" json:"layout,omitempty"`
	User     *User           `gorm:"foreignKey:UserUUID;references:UUID" json:"user,omitempty"`
}

func (f *UserLayout) BeforeCreate(tx *gorm.DB) (err error) {
	if f.UUID == "" {
		f.UUID = nuuid.ShortUUID("lay")
	}
	return
}
