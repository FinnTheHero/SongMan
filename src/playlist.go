package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/zmb3/spotify"
)

/* Generate an exportable track blueprint */
func GenerateTrackBlueprint(client spotify.Client, trackID spotify.ID) TrackBlueprint {

	var blueprint TrackBlueprint

	// Get track
	track := GetTrack(client, trackID)

	// Assign values to the blueprint
	blueprint.TrackID = track.ID
	blueprint.TrackName = track.Name
	blueprint.Artists = track.Artists
	blueprint.Duration = track.Duration
	blueprint.Link = track.ExternalURLs["spotify"]

	return blueprint
}

/* Generate an exportable playlist blueprint */
func GeneratePlaylistBlueprint(client spotify.Client, playlistID spotify.ID) PlaylistBlueprint {

	var blueprint PlaylistBlueprint

	// Get playlist
	playlist := GetPlaylist(client, playlistID)

	// Get tracks
	tracks := GetTracks(playlist)

	// Assign values to the blueprint
	blueprint.PlaylistID = playlist.ID
	blueprint.Description = playlist.Description
	blueprint.Tracks = tracks
	blueprint.Name = playlist.SimplePlaylist.Name

	return blueprint
}

/* Export the Track blueprint */
func ExportTrackBlueprint(blueprint TrackBlueprint) {
	// Create the blueprint directory if it doesn't exist
	blueprintDir := "../blueprints"
	err := os.MkdirAll(blueprintDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Convert blueprint to JSON
	blueprintJSON, err := json.Marshal(blueprint)
	if err != nil {
		panic(err)
	}

	// Write JSON to file
	blueprintFileName := blueprint.TrackName + ".json"
	blueprintFile := filepath.Join(blueprintDir, blueprintFileName)
	err = os.WriteFile(blueprintFile, blueprintJSON, 0644)
	if err != nil {
		panic(err)
	}
}

/* Export the Playlist blueprint */
func ExportPlaylistBlueprint(blueprint PlaylistBlueprint) {
	// Create the blueprint directory if it doesn't exist
	blueprintDir := "../blueprints"
	err := os.MkdirAll(blueprintDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Convert blueprint to JSON
	blueprintJSON, err := json.Marshal(blueprint)
	if err != nil {
		panic(err)
	}

	// Write JSON to file
	blueprintFileName := blueprint.Name + ".json"
	blueprintFile := filepath.Join(blueprintDir, blueprintFileName)
	err = os.WriteFile(blueprintFile, blueprintJSON, 0644)
	if err != nil {
		panic(err)
	}
}

/* Return full playlist details using client and playlist ID. */
func GetPlaylist(client spotify.Client, playlistID spotify.ID) *spotify.FullPlaylist {
	// Get the playlist ID
	p := spotify.ID(playlistID)

	// Get the playlist
	playlist, err := client.GetPlaylist(p)

	if err != nil {
		panic(err)
	}

	return playlist
}

/* Return the list of tracks from the playlist. */
func GetTracks(playlist *spotify.FullPlaylist) spotify.PlaylistTrackPage {
	// Get the playlist tracks
	tracks := playlist.Tracks

	return tracks
}

func GetTrack(client spotify.Client, trackID spotify.ID) spotify.FullTrack {
	// Get the track ID
	t := spotify.ID(trackID)

	// Get the track
	track, err := client.GetTrack(t)
	if err != nil {
		panic(err)
	}

	return *track
}
