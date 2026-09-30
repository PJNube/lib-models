package datatypes

import (
	"fmt"
	"strings"
)

type UserType string

const (
	// UserTypeInteractive is a human account that logs in with credentials.
	// The zero value ("") is treated as Interactive so pre-existing rows and
	// callers that never set the field keep today's behavior.
	UserTypeInteractive UserType = "Interactive"
	// UserTypeCloudRepresentative is the non-interactive Representative
	// Account standing in for one enrolled cloud platform: no password, no
	// login, never a superuser.
	UserTypeCloudRepresentative UserType = "CloudRepresentative"
)

var UserTypeMap = map[UserType]struct{}{
	UserTypeInteractive:         {},
	UserTypeCloudRepresentative: {},
}

func (dt *UserType) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := UserTypeMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(UserTypeMap))
	for m := range UserTypeMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid user type, i.e.: %s", strings.Join(v, " or "))
}

// IsInteractive reports whether the account is a human login account. The
// zero value counts as interactive.
func (dt UserType) IsInteractive() bool {
	return dt == "" || dt == UserTypeInteractive
}
