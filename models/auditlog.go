package models

import (
	"time"

	"gorm.io/datatypes"
)

type AuditLog struct {
	ID           uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	Username     string         `json:"username,omitempty" gorm:"size:128"`
	ClientIP     string         `json:"clientIp,omitempty" gorm:"size:45"`
	Action       string         `json:"action,omitempty" gorm:"not null;size:64"`
	Url          string         `json:"url,omitempty" gorm:"type:text;not null;index:idx_audit_logs_url_created_at"`
	Body         datatypes.JSON `json:"body,omitempty" gorm:"type:jsonb"`
	PreviousBody datatypes.JSON `json:"previousBody,omitempty" gorm:"type:jsonb"`
	CreatedAt    time.Time      `json:"auditTime,omitempty" gorm:"not null;autoCreateTime;index:idx_audit_logs_url_created_at"`

	// Cloud integration fields. Empty on local (web) entries; Source
	// distinguishes them. The actor fields are what the cloud *claimed* —
	// ActorAsserted records that they were never proven.
	Source             string `json:"source,omitempty" gorm:"size:16;default:local"`
	Platform           string `json:"platform,omitempty" gorm:"size:128"`
	RepresentativeUser string `json:"representativeUser,omitempty" gorm:"size:128"`
	OnBehalfOf         string `json:"onBehalfOf,omitempty" gorm:"size:128"`
	AssertedRole       string `json:"assertedRole,omitempty" gorm:"size:128"`
	ActorAsserted      bool   `json:"actorAsserted,omitempty" gorm:"default:false"`
	Outcome            string `json:"outcome,omitempty" gorm:"size:32"`
}
