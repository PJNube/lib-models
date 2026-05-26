package models

import (
	"time"

	"gorm.io/gorm"
)

type UserVisit struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserUUID  string    `json:"userUuid" gorm:"index;not null;uniqueIndex:idx_user_uuid_data"`
	Data      string    `json:"data" gorm:"type:TEXT;not null;uniqueIndex:idx_user_uuid_data"`
	VisitedAt time.Time `json:"visitedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func (v *UserVisit) BeforeCreate(tx *gorm.DB) (err error) {
	if v.VisitedAt.IsZero() {
		v.VisitedAt = time.Now()
	}
	return
}
