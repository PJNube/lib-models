package models

import (
	"github.com/PJNube/lib-models/datatypes"
	"github.com/PJNube/lib-models/utils/nuuid"
	"gorm.io/gorm"
)

type User struct {
	UUID        string               `json:"uuid" sql:"uuid" gorm:"type:varchar(255);unique;primaryKey"`
	Username    string               `json:"username" gorm:"unique;not null"`
	Password    string               `json:"password"`
	Status      datatypes.UserStatus `json:"status" gorm:"default:Disabled"`
	IsSuperuser bool                 `json:"isSuperuser" gorm:"default:false"`

	// Extra fields
	Role string `gorm:"-" json:"role"`
}

func (r *User) BeforeCreate(tx *gorm.DB) (err error) {
	if r.UUID == "" {
		uuid, _ := nuuid.MakeUUID()
		r.UUID = uuid
	}
	return
}
