package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// Load environment variables dari file .env 
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	// Ambil konfigurasi database dari file .env 
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Buat DSN (Data Source Name) utk koneksi ke database MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPass, dbHost, dbPort, dbName)
	
	// Buka koneksi ke database MySQL dengan GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Cek apakah ada error saat membuka koneksi ke database/ error handling
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	return db
}
