package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	const MYSQL = "root:Password#@tcp(127.0.0.1:3306)/learninggo?charset=utf8mb4&parseTime=true&loc=Local"
	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Connected to database")
}
