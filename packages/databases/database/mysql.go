package database

import (
	"fmt"
	"log"
	configs "xyz/packages/databases/config"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL(e *echo.Echo) *gorm.DB {
	config, err := configs.LoadConfig()
	if err != nil {
		e.Logger.Fatalf("failed to load configuration: %v", err)
	}
	// Initialize MySQL connection

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MYSQL.MYSQL_USER,
		config.MYSQL.MYSQL_PASS,
		config.MYSQL.MYSQL_HOST,
		config.MYSQL.MYSQL_PORT,
		config.MYSQL.MYSQL_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect MySQL: %v", err)
	}

	return db
}