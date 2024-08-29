package db

import (
	"fmt"
	"transaction/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectionToDB() (err error) {
	data := settings.Settings.PgSettings

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Dushanbe", data.Host, data.User, data.Password, data.DBName),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return
	}
	return
}
