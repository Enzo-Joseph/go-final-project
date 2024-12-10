package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")

	dsn := user + ":" + pass + "@tcp(127.0.0.1:3306)/classrooms?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database.")
	}

	// Migrate the schema
	db.AutoMigrate(&Lecturer{}, &Student{}, &Course{}, &Post{})
	db.AutoMigrate(&Post{})

	DB = db
}
