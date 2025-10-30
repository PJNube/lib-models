package dtos

type Role struct {
	Name   string   `json:"name"`
	Source string   `json:"source"`
	Topics []string `json:"topics"`
	Effect *string  `json:"effect"`
}
