package twiters

import (
	"time"

	"gorm.io/gorm"
)

type Tweet struct {
	Id        int            `gorm:"primaryKey" json:"id"`
	Tweet     string         `json:"tweet"`
	UserId    int            `json:"-"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
