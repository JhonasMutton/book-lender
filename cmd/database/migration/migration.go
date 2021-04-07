package main

import (
	"fmt"
	"github.com/JhonasMutton/book-lender/pkg/database"
	"github.com/JhonasMutton/book-lender/pkg/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)
func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error to start migration: cannot load environment variables: " + err.Error())
	}
}

func main() {
	println("Starting migration!")
	dsn := fmt.Sprintf(database.MySqlDsnFormat,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}

	if err := db.Debug().AutoMigrate(model.User{}, model.LoanBook{}, model.Book{}); err != nil {
		panic("error to migration:" + err.Error())
	}

	println("Migration was successful")
}
