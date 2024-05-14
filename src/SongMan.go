package main

import (
	"SongMan/blueprint"
	"SongMan/download"
	s "SongMan/spotify"
	"SongMan/utils"
	"flag"
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

	HandleProcessing()
}

func HandleProcessing() {
	client := s.HandleSpotify()

	id, mode := utils.ExtractSpotifyID(Link)

	switch mode {
	case "track":
		track := s.GetTrack(client, id)
		blueprint.ExportTrackBlueprint(track)

		if Download {
			download.DownloadTrack(track)
		}
	case "playlist":
		playlist := s.GetPlaylist(client, id)
		blueprint.ExportPlaylistBlueprint(playlist)

		if Download {
			download.DownloadPlaylist(playlist)
		}
	}
}
