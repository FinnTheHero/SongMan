package download

import (
	"context"
	"fmt"

	"github.com/kkdai/youtube/v2"
	"github.com/kkdai/youtube/v2/downloader"
)

func DownloadTrack(trackName string) {
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

	/* IN CASE FFMPEG IS NOT AVAILABLE - MAY NOT WORK SOMETIMES */
	// fmt.Println("ffmpeg is not installed.")
	// fmt.Println("Ffmpeg is not installed!\n! ! ! Some downloads may not work ! ! !\nPlease isntall ffmpeg and try again.")

	// var format string
	// if strings.Contains(video.Formats[1].MimeType, ";") {
	// 	format = strings.Split(video.Formats[1].MimeType, ";")[0]
	// 	if format == "video/mp4" {
	// 		format = ".mp4"
	// 	} else if format == "video/webm" {
	// 		format = ".webm"
	// 	}
	// }

	// fmt.Println("format:", format)
	// d.Download(ctx, video, &video.Formats[1], trackName+format)

	d.DownloadComposite(ctx, trackName+".mp4", video, "", "mp4", "")
}

/* Download playlist */
// func DownloadPlaylist(playlistName string) {

// }
