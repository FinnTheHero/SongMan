package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func GetWorkinDir() string {
	fileDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return fileDir
}

func AppendTrackDetailsToFile(trackName string, fileName string, dir string) {
	if !CheckFileExistence(fileName, dir) {
		// Create the file
		file, err := os.Create(filepath.Join(dir, fileName))
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	// Read the existing JSON data from the file
	existingData, err := os.ReadFile(filepath.Join(dir, fileName))
	if err != nil {
		panic(err)
	}

	// Parse the existing JSON data into a map
	var jsonData map[string]interface{}
	err = json.Unmarshal(existingData, &jsonData)
	if err != nil {
		panic(err)
	}

	// Create a new track details map
	trackDetails := make(map[string]string)
	trackDetails["TrackName"] = trackName
	// Add other details to the track details map

	// Append the new track details to the "downloaded" array

	downloadedArray, ok := jsonData["downloaded"].([]interface{})
	if !ok {
		downloadedArray = make([]interface{}, 0)
	}
	downloadedArray = append(downloadedArray, trackDetails)
	jsonData["downloaded"] = downloadedArray

	// Convert the updated JSON data back to bytes
	updatedData, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// Write the updated JSON data to the file
	err = os.WriteFile(filepath.Join(dir, fileName), updatedData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Track cached successfully.")
}

// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
