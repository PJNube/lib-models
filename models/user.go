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
	Layouts     []*UserLayout        `json:"layouts,omitempty" gorm:"foreignKey:UserUUID;constraint:OnDelete:CASCADE"`

	// Extra fields
	Roles       []string      `gorm:"-" json:"roles"`
	UserLayouts []*UserLayout `json:"userLayouts,omitempty" gorm:"foreignKey:UserUUID;references:UUID"`
}

func (r *User) BeforeCreate(tx *gorm.DB) (err error) {
	if r.UUID == "" {
		uuid, _ := nuuid.MakeUUID()
		r.UUID = uuid
	}
	return
}
