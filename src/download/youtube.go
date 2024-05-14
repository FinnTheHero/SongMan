package download

import (
	"SongMan/types"
	"SongMan/utils"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func GetVideoId(trackName string) string {
	// Read the JSON file
	filepath := "../blueprints/" + trackName + ".json"
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	// Unmarshal the JSON data
	var track types.TrackBlueprint
	err = json.Unmarshal(data, &track)
	if err != nil {
		panic(err)
	}

	// Access the data from the JSON file
	artist := track.Artists[0].Name

	querry := track.Name + " " + artist + " audio"

	ctx := context.Background()

	yt, err := youtube.NewService(ctx, option.WithAPIKey(utils.GetEnvVariables("YOUTUBE_API_KEY")))
	if err != nil {
		panic(err)
	}

	listItems := []string{"id", "snippet"}

	response, err := yt.Search.List(listItems).Q(querry).MaxResults(1).Do()
	if err != nil {
		panic(err)
	}

	fmt.Println("Video id: ", response.Items[0].Id.VideoId)

	return response.Items[0].Id.VideoId
}
