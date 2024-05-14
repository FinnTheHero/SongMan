package download

import (
	"context"
	"fmt"

	"github.com/kkdai/youtube/v2"
	"github.com/kkdai/youtube/v2/downloader"
)

func DownloadAudio(trackName string) {
	// Get the video information
	videoId := GetVideoId(trackName)

	client := youtube.Client{}

	ctx := context.Background()

	video, err := client.GetVideo(videoId)
	if err != nil {
		fmt.Println("Error getting video: ", err)
		return
	}

	d := downloader.Downloader{Client: client, OutputDir: "../videos"}

	d.Download(ctx, video, &video.Formats[1], trackName+".mp4")
}
