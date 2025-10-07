package dtos

type NotificationPayload struct {
	Event string `json:"event,omitempty"`
	Data  any    `json:"data,omitempty"`
}
