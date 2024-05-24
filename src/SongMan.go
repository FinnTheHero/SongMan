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

	HandleProcessing()
}

func HandleProcessing() {
	client := s.HandleSpotify()

	id, mode := utils.ExtractSpotifyID(Link)

	switch mode {
	case "track":
		track := s.GetTrack(client, id)

		if exists := utils.CheckFileExistence(track.Name+".json", "../blueprints"); exists {
			fmt.Println("'Track' blueprint for '" + track.Name + "' already exists.\n")
		} else {
			blueprint.ExportTrackBlueprint(track)
		}

		if Download {
			utils.CreateDirIfNotExist("../videos")
			utils.CreateDirIfNotExist("../music")
			download.DownloadTrack(track)
			download.ConvertToMp3(track.Name)
			download.A_process(track, ".mp3")
		}

	case "playlist":
		playlist := s.GetPlaylist(client, id)

		if exists := utils.CheckFileExistence(playlist.Name+".json", "../blueprints"); exists {
			fmt.Println("'Playlist' blueprint for '" + playlist.Name + "' already exists.\n")
		} else {
			blueprint.ExportPlaylistBlueprint(playlist)
		}

		if Download {
			utils.CreateDirIfNotExist("../videos")
			utils.CreateDirIfNotExist("../music")
			download.DownloadPlaylist(playlist)
		}
	}
}
