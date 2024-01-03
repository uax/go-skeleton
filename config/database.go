package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := Config("DB_USERNAME") + ":" + Config("DB_PASSWORD") + "@tcp(" + Config("DB_HOST") + ":" + Config("DB_PORT") + ")/" + Config("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
}
