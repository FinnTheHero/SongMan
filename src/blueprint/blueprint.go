package blueprint

import (
	"SongMan/utils"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zmb3/spotify"
)

/* Export the Track blueprint */
func ExportTrackBlueprint(track spotify.FullTrack) {
	// Create the blueprint directory if it doesn't exist
	blueprintDir := "../blueprints"
	err := os.MkdirAll(blueprintDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Convert blueprint to JSON
	blueprintJSON, err := json.Marshal(track)
	if err != nil {
		panic(err)
	}

	// Replace spaces and slashes in the Track name
	// To avoid error while creating the file
	trackName := strings.ReplaceAll(track.Name, "/", "_")
	trackName = strings.ReplaceAll(trackName, " ", "_")

	// Write JSON to file
	blueprintFileName := trackName + ".json"
	blueprintFile := filepath.Join(blueprintDir, blueprintFileName)
	err = os.WriteFile(blueprintFile, blueprintJSON, 0644)
	if err != nil {
		panic(err)
	}

	ExportMessage("Track", trackName)
}

/* Export the Playlist blueprint */
func ExportPlaylistBlueprint(playlist *spotify.FullPlaylist) {
	// Create the blueprint directory if it doesn't exist
	blueprintDir := "../blueprints"
	err := os.MkdirAll(blueprintDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Convert blueprint to JSON
	blueprintJSON, err := json.Marshal(playlist)
	if err != nil {
		panic(err)
	}

	// Replace spaces and slashes in the playlist name
	// To avoid error while creating the file
	playlistName := strings.ReplaceAll(playlist.Name, "/", "_")
	playlistName = strings.ReplaceAll(playlistName, " ", "_")

	// Write JSON to file
	blueprintFileName := playlistName + ".json"
	blueprintFile := filepath.Join(blueprintDir, blueprintFileName)
	err = os.WriteFile(blueprintFile, blueprintJSON, 0644)
	if err != nil {
		panic(err)
	}

	ExportMessage("Playlist", playlistName)
}

/* Export message */
func ExportMessage(mode string, file string) {
	fmt.Println("'" + mode + "' blueprint for '" + file + "' has been exported successfully.")

	fileDir := utils.GetWorkinDir()

	fmt.Println("\nLocation: " + fileDir + "/blueprints/" + file + ".json")
}
