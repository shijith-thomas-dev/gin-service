package config

import (
	"fmt"
	"gin-service/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB, db *gorm.DB

func Connect() {
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open("postgres://user:userP@db:5432/userDB"), &gorm.Config{})
		if err != nil {
			fmt.Printf("Unable to Open DB: %s... Retrying\n", err.Error())
			time.Sleep(time.Second * 10)
		} else {
			fmt.Printf("ERR is nil>>>>>\n")
			err = nil
			break
		}
	}
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Connection to database successful!\n")
	db.AutoMigrate(&models.User{})
	DB = db

}
