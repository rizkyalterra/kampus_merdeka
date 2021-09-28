package users

import (
	"kampus_merdeka/business/users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        int    `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Name      string
	Address   string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *Users) ToDomain() users.Domain {
	return users.Domain{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToListDomain(data []Users) (result []users.Domain) {
	result = []users.Domain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Address:   domain.Address,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
