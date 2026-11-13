package dtos

type EventPayload struct {
	ExtensionID string   `json:"extensionId"` // extensionId must be exist and not empty
	Event       string   `json:"event,omitempty"`
	Payload     any      `json:"payload,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}
