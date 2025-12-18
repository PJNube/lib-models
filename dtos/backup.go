package dtos

type CreateBackupRequest struct {
	Name       string   `json:"name"`
	Components []string `json:"components"`
}
