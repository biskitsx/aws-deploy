package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	var err error
	dsn := os.Getenv("DSN")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		fmt.Println("connect to database fail ??? ")
		return errors.New("can't connect to mysql")
	} else {
		fmt.Println("connect to database successfully !!!")
		return nil
	}
}
