package datatypes

import (
	"fmt"
	"strings"
)

type BackupComponent string

const (
	BackupComponentControlApp BackupComponent = "control-app"
	BackupComponentData       BackupComponent = "data"
	BackupComponentAll        BackupComponent = "all"
)

var BackupComponentMap = map[BackupComponent]struct{}{
	BackupComponentControlApp: {},
	BackupComponentData:       {},
	BackupComponentAll:        {},
}

func (dt *BackupComponent) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := BackupComponentMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(BackupComponentMap))
	for m := range BackupComponentMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid backup component i.e.: %s", strings.Join(v, " or "))
}

type CreateStatus string

const (
	CreateStatusCompleted  CreateStatus = "completed"
	CreateStatusInProgress CreateStatus = "in-progress"
	CreateStatusFailed     CreateStatus = "failed"
)

var CreateStatusMap = map[CreateStatus]struct{}{
	CreateStatusCompleted:  {},
	CreateStatusInProgress: {},
	CreateStatusFailed:     {},
}
