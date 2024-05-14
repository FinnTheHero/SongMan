package blueprint

import (
	"SongMan/utils"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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

	// Write JSON to file
	blueprintFileName := track.Name + ".json"
	blueprintFile := filepath.Join(blueprintDir, blueprintFileName)
	err = os.WriteFile(blueprintFile, blueprintJSON, 0644)
	if err != nil {
		panic(err)
	}

	ExportMessage("Track", track.Name)
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

	// Write JSON to file
	blueprintFileName := playlist.Name + ".json"
	blueprintFile := filepath.Join(blueprintDir, blueprintFileName)
	err = os.WriteFile(blueprintFile, blueprintJSON, 0644)
	if err != nil {
		panic(err)
	}

	ExportMessage("Playlist", playlist.Name)
}

/* Export message */
func ExportMessage(mode string, file string) {
	fmt.Println("'" + mode + "' blueprint for '" + file + "' has been exported successfully.")

	fileDir := utils.GetWorkinDir()

	fmt.Println("\nLocation: " + fileDir + "/blueprints/" + file + ".json")
}
