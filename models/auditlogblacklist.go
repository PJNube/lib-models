package models

import (
	"time"
)

type AuditLogBlacklist struct {
	ID uint64 `json:"id" gorm:"primaryKey;autoIncrement"`

	Pattern     string `json:"pattern" gorm:"type:text;not null;uniqueIndex"`
	Description string `json:"description" gorm:"size:255"`

	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"autoUpdateTime"`
}
