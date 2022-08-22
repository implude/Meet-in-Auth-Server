package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	var dsn = os.Getenv("DB_DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrateTable(database)
	DB = database
}

func migrateTable(database *gorm.DB) {
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Token{})
}
