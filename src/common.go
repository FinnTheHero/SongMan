package main

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariables(key string) string {
	err := godotenv.Load("../.env")

	if err != nil {
		panic("Error loading .env file")
	}

	if key == "" {
		panic("Key is empty")
	} else if key != "CLIENT_ID" && key != "CLIENT_SECRET" {
		panic("Key is invalid")
	}

	return os.Getenv(key)
}
