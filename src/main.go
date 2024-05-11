package main

import (
	"flag"
	"fmt"

	"SongMan/bp"
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
	authConfig := bp.GetConfigCredentials()

	// Create a new client
	client := bp.GetNewUser(authConfig)

	if trackMode == "" && playlistMode == "" {
		flag.PrintDefaults()
		fmt.Println("track link: ", trackMode)
		fmt.Println("playlist link: ", playlistMode)
	} else if playlistMode != "" {
		playlistID := bp.ExtractSpotifyID(playlistMode)
		blueprint := bp.GeneratePlaylistBlueprint(client, playlistID)
		bp.ExportPlaylistBlueprint(blueprint)
	} else if trackMode != "" {
		trackID := bp.ExtractSpotifyID(trackMode)
		blueprint := bp.GenerateTrackBlueprint(client, trackID)
		bp.ExportTrackBlueprint(blueprint)
	}
}
