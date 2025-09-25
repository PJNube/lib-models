package dtos

import (
	"github.com/PJNube/lib-models/enums/notification"
)

type NotificationPayload struct {
	Event notification.Event `json:"event,omitempty"`
	Data  any                `json:"data,omitempty"`
}
