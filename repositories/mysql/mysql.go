package mysql

import (
	"backend-mental-guardians/repositories/mysql/content"
	"backend-mental-guardians/repositories/mysql/music"
	"backend-mental-guardians/repositories/mysql/user"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort string
}

func ConnectDB(config Config) *gorm.DB {
	dbportint, _ := strconv.Atoi(config.DBPort)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		dbportint,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate(db)
	return db
}

func InitMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(user.User{}, music.Music{}, content.Content{}); err != nil {
		log.Println("Error migrating user table")
	}
}