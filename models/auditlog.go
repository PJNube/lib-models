package models

import (
	"time"

	"github.com/PJNube/lib-models/enums"
	"gorm.io/datatypes"
)

type AuditLog struct {
	ID            uint64              `json:"id" gorm:"primaryKey;autoIncrement"`
	Username      string              `json:"username,omitempty" gorm:"size:128"`
	ClientIP      string              `json:"clientIp,omitempty" gorm:"size:45"`
	Action        string              `json:"action,omitempty" gorm:"not null;size:64"`
	Url           string              `json:"url,omitempty" gorm:"not null;size:256"`
	ResourceId    string              `json:"resourceId,omitempty" gorm:"not null"`
	Body          datatypes.JSON      `json:"body,omitempty" gorm:"type:json"`
	PreviousBody  datatypes.JSON      `json:"previousBody,omitempty" gorm:"type:json"`
	CreatedAt     time.Time           `json:"auditTime,omitempty" gorm:"not null;autoCreateTime"`
	OperationType enums.OperationType `json:"operationType,omitempty"`
}
