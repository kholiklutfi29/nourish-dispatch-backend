package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetDatabaseURL() string {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		dbname,
		sslmode,
	)
}

func GetJWTSecret() string {

	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		panic("JWT_SECRET is not configured")
	}

	return secret
}