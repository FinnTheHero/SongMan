package main

import (
	"log"
)

func main() {

	// Configure client credentials
	authConfig := GetConfigCredentials()

	// Create a new client
	client := GetNewUser(authConfig)

	playlist := GetPlaylist(client, "37i9dQZF1DXcBWIGoYBM5M")

	// Print the playlist details
	log.Println("playlist id:", playlist.ID)
	log.Println("playlist name:", playlist.Name)
	log.Println("playlist description:", playlist.Description)
}
