package main

import "github.com/zmb3/spotify"

type PlaylistBlueprint struct {
	PlaylistID  string                    `json:"playlist_id"`
	Description string                    `json:"description"`
	Tracks      spotify.PlaylistTrackPage `json:"tracks"`
	Name        string                    `json:"playlist_name"`
}
