package main

import (
	"flag"

	"SongMan/blueprint"
	"SongMan/download"
	"SongMan/utils"
)

var help bool
var Link string
var Download bool

func init() {
	flag.BoolVar(&help, "help", false, "Display this help message")
	flag.StringVar(&Link, "link", "", "The Spotify link to the track or playlist")
	flag.BoolVar(&Download, "download", false, "Download the provided track or playlist")
}

func main() {

	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	// Configure client credentials
	authConfig := blueprint.GetConfigCredentials()

	// Create a new client
	client := blueprint.GetNewUser(authConfig)

	if Link == "" {
		flag.PrintDefaults()
		return
	}

	id, mode := utils.ExtractSpotifyID(Link)

	switch mode {
	case "track":
		track := utils.GetTrack(client, id)
		blueprint.ExportTrackBlueprint(track)

		if Download {
			download.DownloadTrack(track.Name)
		}
	case "playlist":
		playlist := utils.GetPlaylist(client, id)
		blueprint.ExportPlaylistBlueprint(playlist)

		if Download {
			download.DownloadTrack(playlist.Name)
		}
	}
}
