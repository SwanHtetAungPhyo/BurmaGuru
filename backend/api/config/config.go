package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%v", err.Error())
	}
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GetAccessTokenExpiration() string {
	return os.Getenv("ACCESS_TOKEN_EXPIRATION")
}

func GetRefreshTokenExpiration() string {
	return os.Getenv("REFRESH_TOKEN_EXPIRATION")
}
