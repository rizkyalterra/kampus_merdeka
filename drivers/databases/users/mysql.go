package users

import (
	"context"
	"kampus_merdeka/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Login(ctx context.Context, email string, password string) (users.Domain, error) {
	var user Users
	result := rep.Conn.First(&user, "email = ? AND password = ?",
		email, password)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil

}
