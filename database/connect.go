package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB
var err error

// Подключение к базе данных
// В будущем логично перенести параметры соединения в
// отдельный конфигурационный файл
func Connect() {
	dsn := "gonews:123123@tcp(127.0.0.1:3306)/go_news?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

// Отключение от базы данных
func Disconnect() {
	connection, _ := DB.DB()
	connection.Close()
}