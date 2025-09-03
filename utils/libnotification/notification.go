package libnotification

import (
	"encoding/json"

	"github.com/PJNube/lib-models/dtos"
)

func ParseNotificationData(subject string, data []byte) (*dtos.NotificationPayload, error) {
	var req dtos.APIResponse
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, err
	}

	dataBytes, err := json.Marshal(req.Data)
	if err != nil {
		return nil, err
	}

	var payload dtos.NotificationPayload
	if err := json.Unmarshal(dataBytes, &payload); err != nil {
		return nil, err
	}

	payload.Subject = subject

	return &payload, nil
}
