package main

import (
	"SongMan/blueprint"
	"SongMan/download"
	s "SongMan/spotify"
	"SongMan/utils"
	"flag"
	"fmt"
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

	// Parse and handle arguments
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	if Link == "" {
		flag.PrintDefaults()
		return
	}

	// Create folders if they don't exist
	utils.CreateDirIfNotExist("../blueprints")
	utils.CreateDirIfNotExist("../videos")
	utils.CreateDirIfNotExist("../music")

	// Handle the downloading and blueprinting
	HandleProcessing()
}

func HandleProcessing() {

	client := s.HandleSpotify()

	// Get the ID and mode of the link
	id, mode := s.ExtractSpotifyID(Link)

	// Check if file exists already
	utils.CheckFileExistence(id.String()+".json", "../blueprints")

	switch mode {
	case "track":
		track := s.GetTrack(client, id)

		if utils.CheckFileExistence(track.Name+".json", "../blueprints") {
			fmt.Println("'Track' blueprint for '" + track.Name + "' already exists.\n")
		} else {
			blueprint.ExportTrackBlueprint(track)
		}

		if Download {
			download.DownloadTrack(track, "")
			download.ConvertToMp3(track, "")
		}

	case "playlist":
		playlist := s.GetPlaylist(client, id)

		if utils.CheckFileExistence(playlist.Name+".json", "../blueprints") {
			fmt.Println("'Playlist' blueprint for '" + playlist.Name + "' already exists.\n")
		} else {
			blueprint.ExportPlaylistBlueprint(playlist)
		}

		if Download {
			download.DownloadPlaylist(playlist)
		}
	}
}
