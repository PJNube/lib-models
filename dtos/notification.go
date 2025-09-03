package dtos

import "github.com/PJNube/lib-models/datatypes"

type NotificationPayload struct {
	Event       datatypes.NotificationEvent `json:"event"`
	Body        any                         `json:"body"`
	Source      string                      `json:"source"`
	Subject     string                      `json:"subject"`
	ExtensionId string                      `json:"extensionId"`
}
