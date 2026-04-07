package models

import (
	"github.com/PJNube/lib-models/utils/nuuid"
	"github.com/PJNube/lib-models/utils/pjnjson"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserPreference struct {
	UUID      string         `json:"uuid" gorm:"primaryKey;type:VARCHAR(255)" `
	UserUUID  *string        `json:"userUuid,omitempty" gorm:"uniqueIndex" `
	Data      datatypes.JSON `json:"data"`
	IsDefault bool           `json:"isDefault" gorm:"default:false"`

	User *User `json:"user,omitempty" gorm:"foreignKey:UserUUID"`
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
