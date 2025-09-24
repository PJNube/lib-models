package dtos

import (
	"github.com/PJNube/lib-models/enums/notification"
)

type NotificationPayload struct {
	Event notification.Event `json:"event,omitempty"`
	Body  any                `json:"body,omitempty"`
}
