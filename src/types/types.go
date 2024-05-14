package types

import "github.com/zmb3/spotify"

type PlaylistBlueprint struct {
	ID          spotify.ID                `json:"playlist_id"`
	Description string                    `json:"description"`
	Tracks      spotify.PlaylistTrackPage `json:"tracks"`
	Name        string                    `json:"playlist_name"`
}

type TrackBlueprint struct {
	ID       spotify.ID             `json:"track_id"`
	Name     string                 `json:"track_name"`
	Artists  []spotify.SimpleArtist `json:"artists"`
	Duration int                    `json:"duration"`
	Link     string                 `json:"link"`
}
