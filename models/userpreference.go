package models

import (
	"github.com/PJNube/lib-models/utils/nuuid"
	"github.com/PJNube/lib-models/utils/pjnjson"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserPreference struct {
	UUID      string         `gorm:"primaryKey;type:VARCHAR(255)" json:"uuid"`
	UserUUID  *string        `gorm:"uniqueIndex:idx_user_preferences_user_uuid" json:"userUuid,omitempty"`
	Data      datatypes.JSON `json:"data"`
	IsDefault bool           `json:"isDefault"`

	User *User `gorm:"foreignKey:UserUUID" json:"user,omitempty"`
}

func (f *UserPreference) BeforeCreate(tx *gorm.DB) (err error) {
	if f.UUID == "" {
		f.UUID = nuuid.ShortUUID("upr")
	}
	f.Data = pjnjson.MarshalJson(f.Data)
	return
}

func (f *UserPreference) BeforeUpdate(tx *gorm.DB) (err error) {
	f.Data = pjnjson.MarshalJson(f.Data)
	return
}
