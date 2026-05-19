package datatypes

import (
	"fmt"
	"strings"
)

type BackupTier string

const (
	BackupTierControlApp     BackupTier = "control-app"
	BackupTierExtensionsData BackupTier = "extensions-data"
	BackupTierAll            BackupTier = "all"
)

var BackupTierMap = map[BackupTier]struct{}{
	BackupTierControlApp:     {},
	BackupTierExtensionsData: {},
	BackupTierAll:            {},
}

func (dt *BackupTier) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := BackupTierMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(BackupTierMap))
	for m := range BackupTierMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid backup tier i.e.: %s", strings.Join(v, " or "))
}
