package dtos

type Resource struct {
	Id          string   `json:"id"`
	Description string   `json:"description"`
	Items       []string `json:"items"`
}
