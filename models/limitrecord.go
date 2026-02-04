package models

type LimitRecord struct {
	Key          string `json:"key" gorm:"primaryKey"`
	FailCount    int    `json:"failCount" gorm:"default:0"`
	LockMinutes  int    `json:"lockMinutes" gorm:"default:0"`
	LastFailTime int64  `json:"lastFailTime" gorm:"default:0"`
}
