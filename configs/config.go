package configs

import (
	"backend-mental-guardians/repositories/mysql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default value")
	}
}

func InitConfigMySQL() mysql.Config {
	return mysql.Config{
		DBName: os.Getenv("DBName"),
		DBUser: os.Getenv("DBUser"),
		DBPass: os.Getenv("DBPass"),
		DBHost: os.Getenv("DBHost"),
		DBPort: os.Getenv("DBPort"),
	}
}

func InitConfigJWT() string {
	return os.Getenv("SECRET_JWT")
}

func InitConfigCloudinary() string {
	return os.Getenv("CLOUDINARY_URL")
}

func InitConfigKeyChatbot() string {
	return os.Getenv("AI_KEY")
}
