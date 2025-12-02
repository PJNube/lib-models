package models

import (
	"time"

	"github.com/PJNube/lib-models/utils/nuuid"
	"github.com/PJNube/lib-models/utils/pjnjson"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AuditLog struct {
	UUID               string         `json:"uuid" gorm:"primaryKey;type:VARCHAR(255)"`
	UserUUID           string         `json:"userUuid,omitempty"`
	Username           string         `json:"username,omitempty"`
	AuditTime          time.Time      `json:"auditTime,omitempty" gorm:"not null;autoCreateTime"`
	UserAgent          string         `json:"userAgent,omitempty"`
	Action             string         `json:"action,omitempty"`
	RequestUrl         string         `json:"requestUrl,omitempty"`
	RequestQueryParams string         `json:"requestQueryParams,omitempty"`
	RequestBody        datatypes.JSON `json:"requestBody,omitempty"`
}

func (f *AuditLog) BeforeCreate(tx *gorm.DB) (err error) {
	if f.UUID == "" {
		f.UUID = nuuid.ShortUUID("aud")
	}
	f.RequestBody = pjnjson.MarshalJson(f.RequestBody)
	return
}

func (f *AuditLog) BeforeUpdate(tx *gorm.DB) (err error) {
	f.RequestBody = pjnjson.MarshalJson(f.RequestBody)
	return
}
