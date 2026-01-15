package dtos

import (
	"github.com/PJNube/lib-models/datatypes"
)

type User struct {
	UUID               string               `json:"uuid"`
	Username           string               `json:"username"`
	PasswordExpiry     *int64               `json:"passwordExpiry"` // Unix timestamp in seconds
	PasswordExpiryDays *int                 `json:"passwordExpiryDays"`
	Status             datatypes.UserStatus `json:"status"`

	Roles []*Role `json:"roles,omitempty"`
}
