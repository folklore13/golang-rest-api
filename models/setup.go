package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectToDatabase(){
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/golang-rest-api"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	DB = db
}