package dtos

import "github.com/pjnube/lib-models/datatypes"

type User struct {
	UUID     string               `json:"uuid"`
	Username string               `json:"username"`
	Password string               `json:"password"`
	Status   datatypes.UserStatus `json:"status"`
	Roles    []*Role              `json:"roles"`
}
