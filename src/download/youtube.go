package download

import (
	"SongMan/utils"
	"context"
	"fmt"
	"strings"

	"github.com/zmb3/spotify"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

/* Find and return ID of corresponding youtube video */
func GetVideoId(track spotify.FullTrack) string {
	var isrc string = ""
	var querry string

	if track.ExternalIDs["isrc"] != "" {
		fmt.Println("ISRC - provided: ", track.ExternalIDs["isrc"])
		querry = track.ExternalIDs["isrc"]
		isrc = querry
	} else {
		fmt.Println("ISRC - NOT provided")

		// Access the data from the JSON file
		artist := track.Artists[0].Name

		querry = track.Name + " " + artist + " audio"
	}

	video := GetVideo(querry, 1)
	if video == nil {
		return ""
	}

	if !strings.Contains(video.Items[0].Snippet.Title, track.Name) {
		if isrc != "" {
			video = GetVideo(track.Name+" "+track.Artists[0].Name+" audio", 2)
			if video == nil {
				return ""
			}
		} else {
			fmt.Println("No video found")
			return ""
		}
	}

	fmt.Println("Video ID - ", video.Items[0].Id.VideoId)
	return video.Items[0].Id.VideoId
}

func GetVideo(querry string, api int) *youtube.SearchListResponse {
	ctx := context.Background()

	api_variable := "YOUTUBE_API_KEY_" + fmt.Sprint(api)

	yt, err := youtube.NewService(ctx, option.WithAPIKey(utils.GetEnvVariables(api_variable)))
	if err != nil {
		panic(err)
	}

	listItems := []string{"id", "snippet"}

	var response *youtube.SearchListResponse

	response, err = yt.Search.List(listItems).Q(querry).MaxResults(1).Do()
	if err != nil {
		panic(err)
	}

	if len(response.Items) == 0 {
		fmt.Println("No video found")
		return nil
	}

	return response
}
