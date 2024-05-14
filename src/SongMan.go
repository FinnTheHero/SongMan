package main

import (
	"flag"

	"SongMan/blueprint"
	"SongMan/utils"
)

var help bool
var Link string

func init() {
	flag.BoolVar(&help, "help", false, "Display the help message")
	flag.StringVar(&Link, "link", "", "The link to the track or playlist")
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

	if Link == "" {
		flag.PrintDefaults()
	} else if Link != "" {
		id, mode := utils.ExtractSpotifyID(Link)
		blueprint.GenerateBlueprint(client, id, mode)
	}
}
