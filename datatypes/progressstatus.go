package datatypes

import (
	"fmt"
	"strings"
)

type ProgressStatus string

const (
	ProgressStatusCompleted  ProgressStatus = "completed"
	ProgressStatusInProgress ProgressStatus = "in-progress"
	ProgressStatusFailed     ProgressStatus = "failed"
)

var ProgressStatusMap = map[ProgressStatus]struct{}{
	ProgressStatusCompleted:  {},
	ProgressStatusInProgress: {},
	ProgressStatusFailed:     {},
}

func (dt *ProgressStatus) Validate() error {
	if dt == nil || *dt == "" {
		return nil
	}
	if _, ok := ProgressStatusMap[*dt]; ok {
		return nil
	}
	v := make([]string, 0, len(ProgressStatusMap))
	for m := range ProgressStatusMap {
		v = append(v, string(m))
	}
	return fmt.Errorf("please provide a valid progress status i.e.: %s", strings.Join(v, " or "))
}
