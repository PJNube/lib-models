package models

import (
	"encoding/json"

	"github.com/PJNube/lib-models/utils/nuuid"
	"gorm.io/gorm"
)

type UserLayout struct {
	UUID     string          `gorm:"primaryKey;type:VARCHAR(255)" json:"uuid"`
	Name     string          `gorm:"not null;uniqueIndex:idx_user_layouts_name_user_uuid" json:"name"`
	UserUUID string          `gorm:"not null;uniqueIndex:idx_user_layouts_name_user_uuid" json:"userUuid,omitempty"`
	Data     json.RawMessage `json:"data,omitempty"`

	User *User `gorm:"foreignKey:UserUUID" json:"user,omitempty"`
}

func (f *UserLayout) BeforeCreate(tx *gorm.DB) (err error) {
	if f.UUID == "" {
		f.UUID = nuuid.ShortUUID("ula")
	}
	return
}
