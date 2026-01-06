package dtos

import (
	"github.com/PJNube/lib-models/datatypes"
	"github.com/PJNube/lib-models/models"
)

type User struct {
	UUID     string               `json:"uuid"`
	Username string               `json:"username"`
	Password string               `json:"password"`
	Status   datatypes.UserStatus `json:"status"`
	Layouts  []*models.UserLayout `json:"layouts"`

	Roles []*Role `json:"roles,omitempty"`
}

type ChangePasswordReq struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
