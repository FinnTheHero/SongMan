package utils

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
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

/* Extracts the Spotify track or playlist ID and mode from a given link */
func ExtractSpotifyID(link string) (spotify.ID, string) {
	// Regular expression patterns for track and playlist links
	trackPattern := regexp.MustCompile(`^https?:\/\/open.spotify.com\/track\/([a-zA-Z0-9]+)`)
	playlistPattern := regexp.MustCompile(`^https?:\/\/open.spotify.com\/playlist\/([a-zA-Z0-9]+)`)

	if trackPattern.MatchString(link) {
		// Extract the track ID from the link
		matches := trackPattern.FindStringSubmatch(link)
		return spotify.ID(matches[1]), "track"
	} else if playlistPattern.MatchString(link) {
		// Extract the playlist ID from the link
		matches := playlistPattern.FindStringSubmatch(link)
		return spotify.ID(matches[1]), "playlist"
	} else {
		// Return an empty ID if the link is not a valid track or playlist link
		panic("Invalid link")
	}
}

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

func GetWorkinDir() string {
	fileDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return fileDir
}
