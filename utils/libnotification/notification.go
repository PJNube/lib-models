package libnotification

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PJNube/lib-models/dtos"
	"github.com/PJNube/lib-models/libconstants"
)

func ParseNotificationData(data []byte) (*dtos.NotificationPayload, error) {
	var payload dtos.NotificationPayload
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}

func GenerateNotificationTopic(profile, vendor string, extName ...string) string {
	if len(extName) == 0 {
		return strings.ToLower(fmt.Sprintf("local.%s.%s.%s.$sys", libconstants.EventKeyword, profile, vendor))
	}
	return strings.ToLower(fmt.Sprintf("local.%s.%s.%s.%s", libconstants.EventKeyword, profile, vendor, extName[0]))
}
