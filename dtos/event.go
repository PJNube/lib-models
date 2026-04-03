package dtos

type EventPayload struct {
	ExtensionID string   `json:"extensionId,omitempty"`
	Event       string   `json:"event,omitempty"`
	Payload     any      `json:"payload,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}
