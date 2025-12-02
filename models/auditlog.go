package models

import (
	"time"

	"github.com/PJNube/lib-models/utils/pjnjson"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AuditLog struct {
	ID          uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	Username    string         `json:"username,omitempty" gorm:"size:128"`
	ClientIP    string         `json:"clientIp,omitempty" gorm:"size:45"`
	UserAgent   string         `json:"userAgent,omitempty"`
	Action      string         `json:"action,omitempty" gorm:"not null;size:64"`
	RequestUrl  string         `json:"requestUrl,omitempty" gorm:"not null"`
	RequestBody datatypes.JSON `json:"requestBody,omitempty" gorm:"type:json"`
	CreatedAt   time.Time      `json:"auditTime,omitempty" gorm:"not null;autoCreateTime"`
}

func (f *AuditLog) BeforeCreate(tx *gorm.DB) (err error) {
	f.RequestBody = pjnjson.MarshalJson(f.RequestBody)
	return
}

func (f *AuditLog) BeforeUpdate(tx *gorm.DB) (err error) {
	f.RequestBody = pjnjson.MarshalJson(f.RequestBody)
	return
}
