package models

import (
	"github.com/PJNube/lib-models/utils/nuuid"
	"github.com/PJNube/lib-models/utils/pjnjson"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserLayout struct {
	UUID     string         `gorm:"primaryKey;type:VARCHAR(255)" json:"uuid"`
	Name     string         `gorm:"not null;uniqueIndex:idx_user_layouts_name_user_uuid" json:"name"`
	UserUUID string         `gorm:"not null;uniqueIndex:idx_user_layouts_name_user_uuid" json:"userUuid,omitempty"`
	Data     datatypes.JSON `json:"data,omitempty"`

	User *User `gorm:"foreignKey:UserUUID" json:"user,omitempty"`
}

func (f *UserLayout) BeforeCreate(tx *gorm.DB) (err error) {
	if f.UUID == "" {
		f.UUID = nuuid.ShortUUID("ula")
	}
	f.Data = pjnjson.MarshalJson(f.Data)
	return
}

func (f *UserLayout) BeforeUpdate(tx *gorm.DB) (err error) {
	f.Data = pjnjson.MarshalJson(f.Data)
	return
}
