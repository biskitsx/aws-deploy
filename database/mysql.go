package database

import (
	"fmt"
	"log"
	"os"

	"github.com/biskitsx/aws-deploy/model"
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
		log.Fatal("Error to connect to database")
		return err
	} else {
		Db.AutoMigrate(&model.Attraction{})
		fmt.Println("connect to database successfully !!!")
		return nil
	}
}
