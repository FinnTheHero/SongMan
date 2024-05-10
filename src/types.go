package main

import "github.com/zmb3/spotify"

type PlaylistBlueprint struct {
	PlaylistID  spotify.ID                `json:"playlist_id"`
	Description string                    `json:"description"`
	Tracks      spotify.PlaylistTrackPage `json:"tracks"`
	Name        string                    `json:"playlist_name"`
}

type TrackBlueprint struct {
	TrackID   spotify.ID             `json:"track_id"`
	TrackName string                 `json:"track_name"`
	Artists   []spotify.SimpleArtist `json:"artists"`
	Duration  int                    `json:"duration"`
	Link      string                 `json:"link"`
}