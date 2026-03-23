package dtos

type NotificationPayload struct {
	ExtensionID string   `json:"extensionId,omitempty"`
	Type        string   `json:"type"` // notification, event
	Event       string   `json:"event,omitempty"`
	Payload     any      `json:"payload,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}
