package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseCOnnection() *gorm.DB {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(".env file can't load")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOSTNAME")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSL := os.Getenv("DB_SSL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbPort, dbName, dbSSL)
	db, errr := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errr != nil {
		panic("Failed to create a connection to database")
		//return nil
	}

	return db
}

func DbDisconection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	fmt.Println("close connection from db")
	dbSQL.Close()
}
