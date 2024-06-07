package download

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"SongMan/utils"
	yt "SongMan/youtube"

	"github.com/kkdai/youtube/v2"

	"github.com/kkdai/youtube/v2/downloader"
	"github.com/zmb3/spotify"
)

/* Download single track */
func DownloadTrack(track spotify.FullTrack, playlistName string) {

	utils.Divider()

	// Check if the file already exists
	if ok := utils.CheckFileExistence(track.Name+".mp4", "../videos"); ok {
		fmt.Println("MP4 already exists.")
		return
	}

	// Get the video information
	videoId := yt.GetVideoId(track)
	if videoId == "" {
		return
	}

	client := youtube.Client{}

	ctx := context.Background()

	video, err := client.GetVideo(videoId)
	if err != nil {
		fmt.Println("Error getting video: ", err)
		return
	}

	d := downloader.Downloader{Client: client, OutputDir: "../videos"}

	/* IN CASE FFMPEG IS NOT AVAILABLE - THIS MAY NOT WORK SOMETIMES */
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

	d.DownloadComposite(ctx, track.Name+".mp4", video, "720p", "mp4", "")
	V_process(track, ".mp4")
}

/* Download playlist */
func DownloadPlaylist(playlist *spotify.FullPlaylist) {
	// Loop over playlist
	tracks := playlist.Tracks.Tracks

	for _, track := range tracks {
		DownloadTrack(track.Track, playlist.Name)
		ConvertToMp3(track.Track, playlist.Name)
	}
}

// downloadFile will download a url to a local file.
func DownloadImage(filepath string, url string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
