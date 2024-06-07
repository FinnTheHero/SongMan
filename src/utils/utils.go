package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

/* Retrieve the environment variables from ".env" file. */
func GetEnvVariables(key string) string {
	err := godotenv.Load("../.env")

	if err != nil {
		panic("Error loading .env file")
	}

	if key == "" {
		panic("Key is empty")
	} else if key != "CLIENT_ID" && key != "CLIENT_SECRET" && key != "YOUTUBE_API_KEY_1" && key != "YOUTUBE_API_KEY_2" {
		panic("Key is invalid")
	}

	return os.Getenv(key)
}

/* Check if file exists or not */
func CheckFileExistence(file string, dir string) bool {
	filePath := filepath.Join(dir, file)
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

/* Create a directory if it doesn't exist */
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Directory '" + dir + "' created successfully.")
		}
	}
}

/* Get the working directory */
func GetWorkinDir() string {
	fileDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return fileDir
}

func Divider() {
	fmt.Println("--------------------------------------------------")
}
