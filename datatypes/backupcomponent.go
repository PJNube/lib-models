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
	// BackupTierCustom holds a user-chosen set of components, listed in Backup.Components.
	BackupTierCustom BackupTier = "custom"
	// Firmware and bundle are restore-only tiers, excluded from BackupTierMap/Validate.
	// Their zips are built by hand ({"tier":"firmware"} or {"tier":"bundle"} in the
	// zip comment); the system never produces them. A bundle carries extension
	// packages and optionally firmware.
	BackupTierFirmware BackupTier = "firmware"
	BackupTierBundle   BackupTier = "bundle"
)

var BackupTierMap = map[BackupTier]struct{}{
	BackupTierControlApp:     {},
	BackupTierExtensionsData: {},
	BackupTierAll:            {},
	BackupTierCustom:         {},
}

// BackupComponent is one part of the device's data a backup can hold. Each tier
// is a fixed set of components; a custom backup picks its own.
type BackupComponent string

const (
	BackupComponentControlData   BackupComponent = "control-data"
	BackupComponentDashboard     BackupComponent = "dashboard"
	BackupComponentExtensionData BackupComponent = "extension-data"
	BackupComponentExtensions    BackupComponent = "extensions"
)

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
