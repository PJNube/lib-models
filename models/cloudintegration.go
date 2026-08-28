package models

import (
	"time"

	"github.com/PJNube/lib-models/utils/nuuid"
	"gorm.io/gorm"
)

// CloudIntegration is one enrolled external platform: its pinned public
// key set (Trust Anchor), the Representative Account standing in for it,
// and the NATS connection its requests arrive on. The account's role is
// the Ceiling Role — the most any request from this platform may do.
type CloudIntegration struct {
	UUID           string    `json:"uuid" gorm:"type:varchar(255);unique;primaryKey"`
	Name           string    `json:"name" gorm:"type:varchar(255);unique;not null"` // issuer (iss) of its signed requests
	JWKS           string    `json:"jwks" gorm:"type:text;not null"`                // Trust Anchor: JWKS document as JSON
	UserUUID       string    `json:"userUuid" gorm:"type:varchar(255);not null"`    // Representative Account
	ConnectionUUID string    `json:"connectionUuid" gorm:"type:varchar(255);not null;default:cloud"`
	Enabled        bool      `json:"enabled" gorm:"default:true"`
	CreatedAt      time.Time `json:"createdAt"`
}

func (c *CloudIntegration) BeforeCreate(tx *gorm.DB) (err error) {
	if c.UUID == "" {
		c.UUID = nuuid.ShortUUID("cli")
	}
	return
}
