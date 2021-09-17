package twiters

import (
	"kampus_merdeka/models/users"
	"time"

	"gorm.io/gorm"
)

type Tweet struct {
	Id        int            `gorm:"primaryKey" json:"id"`
	Tweet     string         `json:"tweet"`
	UserId    int            `json:"-"`
	User      users.User     `json:"user" gorm:"foreignKey:UserId"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
