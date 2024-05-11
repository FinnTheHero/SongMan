package main

import (
	"flag"
	"fmt"

	"SongMan/blueprint"
	"SongMan/utils"
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
	authConfig := blueprint.GetConfigCredentials()

	// Create a new client
	client := blueprint.GetNewUser(authConfig)

	if trackMode == "" && playlistMode == "" {
		flag.PrintDefaults()
		fmt.Println("track link: ", trackMode)
		fmt.Println("playlist link: ", playlistMode)
	} else if playlistMode != "" {
		playlistID := utils.ExtractSpotifyID(playlistMode)
		bp := blueprint.GeneratePlaylistBlueprint(client, playlistID)
		blueprint.ExportPlaylistBlueprint(bp)
	} else if trackMode != "" {
		trackID := utils.ExtractSpotifyID(trackMode)
		bp := blueprint.GenerateTrackBlueprint(client, trackID)
		blueprint.ExportTrackBlueprint(bp)
	}
}
