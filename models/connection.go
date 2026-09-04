package models

import (
	"time"

	"github.com/PJNube/lib-models/datatypes"
	"github.com/PJNube/lib-models/utils/nuuid"
	"gorm.io/gorm"
)

// Connection is one broker the device talks to. Config and Credentials are
// JSON blobs whose shape depends on Kind, so a second transport is a new
// Kind rather than a schema change. Credentials are secret: never
// serialized, never returned by any API.
//
// The Local row is a system row: the device writes it from its own config
// at every boot and the API refuses to edit or delete it.
type Connection struct {
	UUID        string                    `json:"uuid" gorm:"type:varchar(255);unique;primaryKey"`
	Name        string                    `json:"name" gorm:"type:varchar(255);unique;not null"`
	Scope       datatypes.ConnectionScope `json:"scope" gorm:"type:varchar(16);not null;unique"`
	Kind        datatypes.ConnectionKind  `json:"kind" gorm:"type:varchar(16);not null"`
	Config      string                    `json:"-" gorm:"type:text;not null"` // kind-specific JSON, non-secret; exposed parsed by the API
	Credentials string                    `json:"-" gorm:"type:text"`          // kind-specific JSON, secret
	System      bool                      `json:"system" gorm:"default:false"`
	Enabled     bool                      `json:"enabled" gorm:"default:true"`
	CreatedAt   time.Time                 `json:"createdAt"`
	UpdatedAt   time.Time                 `json:"updatedAt"`
}

func (c *Connection) BeforeCreate(tx *gorm.DB) (err error) {
	if c.UUID == "" {
		c.UUID = nuuid.ShortUUID("con")
	}
	return
}
