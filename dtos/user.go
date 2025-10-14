package dtos

import (
	"github.com/PJNube/lib-models/datatypes"
	"github.com/PJNube/lib-models/models"
)

type User struct {
	UUID        string               `json:"uuid"`
	Username    string               `json:"username"`
	Password    string               `json:"password"`
	Status      datatypes.UserStatus `json:"status"`
	Roles       []*Role              `json:"roles"`
	UserLayouts []*models.UserLayout `json:"userLayouts"`
}
