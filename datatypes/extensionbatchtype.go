package datatypes

import (
	"fmt"
	"strings"
)

type ExtensionBatchType string

const (
	ExtensionBatchTypeInstall   ExtensionBatchType = "Install"
	ExtensionBatchTypeUninstall ExtensionBatchType = "Uninstall"
)

var ExtensionBatchTypeMap = map[ExtensionBatchType]struct{}{
	ExtensionBatchTypeInstall:   {},
	ExtensionBatchTypeUninstall: {},
}

func (dt *ExtensionBatchType) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := ExtensionBatchTypeMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(ExtensionBatchTypeMap))
	for m := range ExtensionBatchTypeMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid extension batch type, i.e.: %s", strings.Join(v, " or "))
}
