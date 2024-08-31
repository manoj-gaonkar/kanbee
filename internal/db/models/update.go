package models

import (
	"time"
)

type Update struct {
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	Message        string    `gorm:"type:text"`
	Filename       string    `gorm:"type:varchar(255)"`
	AttachmentData []byte    `gorm:"type:blob"`
	ID             uint      `gorm:"primaryKey"`
	TaskID         uint      `gorm:"not null"`
}
