package main

import (
	"flag"
	"fmt"
)

var help bool
var trackMode string
var playlistMode string

func init() {
	flag.BoolVar(&help, "help", false, "Display the help message")
	flag.StringVar(&trackMode, "track", "", "The track link")
	flag.StringVar(&playlistMode, "playlist", "", "The playlist link")
}

func main() {

	flag.Parse()

	if help {
		flag.PrintDefaults()
	}

	// Configure client credentials
	authConfig := GetConfigCredentials()

	// Create a new client
	client := GetNewUser(authConfig)

	if trackMode == "" && playlistMode == "" {
		flag.PrintDefaults()
		fmt.Println("track link: ", trackMode)
		fmt.Println("playlist link: ", playlistMode)
	} else if playlistMode != "" {
		playlistID := ExtractSpotifyID(playlistMode)
		blueprint := GeneratePlaylistBlueprint(client, playlistID)
		ExportPlaylistBlueprint(blueprint)
	} else if trackMode != "" {
		trackID := ExtractSpotifyID(trackMode)
		blueprint := GenerateTrackBlueprint(client, trackID)
		ExportTrackBlueprint(blueprint)
	}
}
