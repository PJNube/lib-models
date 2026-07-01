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
	// BackupTierFirmware is a restore-only tier. Firmware zips are created
	// manually (with {"tier":"firmware"} in the zip comment); the system never
	// produces them, so it is intentionally excluded from BackupTierMap/Validate.
	BackupTierFirmware BackupTier = "firmware"
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
