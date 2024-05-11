package blueprint

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"SongMan/types"
	"SongMan/utils"

	"github.com/zmb3/spotify"
)

/* Generate an exportable track blueprint */
func GenerateTrackBlueprint(client spotify.Client, trackID spotify.ID) types.TrackBlueprint {

	var blueprint types.TrackBlueprint

	// Get track
	track := utils.GetTrack(client, trackID)

	// Assign values to the blueprint
	blueprint.TrackID = track.ID
	blueprint.Name = track.Name
	blueprint.Artists = track.Artists
	blueprint.Duration = track.Duration
	blueprint.Link = track.ExternalURLs["spotify"]

	return blueprint
}

/* Generate an exportable playlist blueprint */
func GeneratePlaylistBlueprint(client spotify.Client, playlistID spotify.ID) types.PlaylistBlueprint {

	var blueprint types.PlaylistBlueprint

	// Get playlist
	playlist := utils.GetPlaylist(client, playlistID)

	// Get tracks
	tracks := utils.GetTracks(playlist)

	// Assign values to the blueprint
	blueprint.PlaylistID = playlist.ID
	blueprint.Description = playlist.Description
	blueprint.Tracks = tracks
	blueprint.Name = playlist.SimplePlaylist.Name

	return blueprint
}

/* Export the Track blueprint */
func ExportTrackBlueprint(blueprint types.TrackBlueprint) {
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

	ExportMessage("ok", "track", blueprint.Name)
}

/* Export the Playlist blueprint */
func ExportPlaylistBlueprint(blueprint types.PlaylistBlueprint) {
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

	ExportMessage("ok", "playlist", blueprint.Name)
}

func ExportMessage(status string, exportType string, name string) {
	if status == "ok" {
		fmt.Printf("Exported '%s'\nBlueprint: '%s'\n", strings.ToUpper(exportType), name)
	} else {
		fmt.Printf("Failed to export '%s'\nBlueprint: '%s'\n", strings.ToUpper(exportType), name)
	}
}
