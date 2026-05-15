package datatypes

import (
	"fmt"
	"strings"
)

type BackupTier string

const (
	BackupComponentControlApp BackupTier = "control-app"
	BackupComponentData       BackupTier = "data"
	BackupComponentAll        BackupTier = "all"
)

var BackupComponentMap = map[BackupTier]struct{}{
	BackupComponentControlApp: {},
	BackupComponentData:       {},
	BackupComponentAll:        {},
}

func (dt *BackupTier) Validate() error {
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
