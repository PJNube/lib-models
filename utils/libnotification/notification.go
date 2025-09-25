package libnotification

import (
	"encoding/json"

	"github.com/PJNube/lib-models/dtos"
)

func ParseNotificationData(data []byte) (*dtos.NotificationPayload, error) {
	var payload dtos.NotificationPayload
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
