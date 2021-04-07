package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

const (
	MySqlVersion   = "8.0.23"
	MySqlDsnFormat = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=60s&readTimeout=60s"
)

func NewDatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf(MySqlDsnFormat,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}

	return db
}
