package dtos

import (
	"github.com/PJNube/lib-models/datatypes"
	"github.com/PJNube/lib-models/models"
)

type User struct {
	UUID           string               `json:"uuid"`
	Username       string               `json:"username"`
	PasswordExpiry *int64               `json:"passwordExpiry"` // Unix timestamp in seconds
	Status         datatypes.UserStatus `json:"status"`
	Layouts        []*models.UserLayout `json:"layouts"`

	Roles []*Role `json:"roles,omitempty"`
}
