package main

import (
	"os"

	"github.com/joho/godotenv"
)

/*
Retrieve the environment variables from ".env" file.

Options: "CLIENT_ID", "CLIENT_SECRET"
*/
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
