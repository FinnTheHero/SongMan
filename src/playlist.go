package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/zmb3/spotify"
)

/* Generate an exportable playlist blueprint */
func GeneratePlaylistBlueprint(client spotify.Client, playlistIDString spotify.ID) PlaylistBlueprint {

	var blueprint PlaylistBlueprint

	// Get playlist
	playlist := GetPlaylist(client, playlistIDString)

	// Get tracks
	tracks := GetTracks(playlist)

	// Assign values to the blueprint
	blueprint.PlaylistID = playlist.ID.String()
	blueprint.Description = playlist.Description
	blueprint.Tracks = tracks
	blueprint.Name = playlist.SimplePlaylist.Name

	return blueprint
}

/* Export the blueprint */
func ExportBlueprint(blueprint PlaylistBlueprint) {
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
func GetPlaylist(client spotify.Client, playlistIDString spotify.ID) *spotify.FullPlaylist {
	// Get the playlist ID
	playlistID := spotify.ID(playlistIDString)

	// Get the playlist
	playlist, err := client.GetPlaylist(playlistID)

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
