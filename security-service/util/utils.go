package util

import (
	"github.com/joho/godotenv"
	"log"
	"golang.org/x/crypto/bcrypt"

)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ComparePassowrd(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func HashPassword(plainPassword string) (string, error) {	
	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassowrd), nil
}