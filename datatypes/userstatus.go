package datatypes

import (
	"fmt"
	"strings"
)

type UserStatus string

const (
	UserStatusEnabled  UserStatus = "Enabled"
	UserStatusDisabled UserStatus = "Disabled"
)

var UserStatusMap = map[UserStatus]struct{}{
	UserStatusEnabled:  {},
	UserStatusDisabled: {},
}

func (dt *UserStatus) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := UserStatusMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(UserStatusMap))
	for m := range UserStatusMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid user status, i.e.: %s", strings.Join(v, " or "))
}
