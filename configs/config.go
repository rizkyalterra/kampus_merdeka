package configs

import (
	"kampus_merdeka/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/kampus_merdeka?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB failed connect")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&users.User{})
}
