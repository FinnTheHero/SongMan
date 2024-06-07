package download

import (
	"SongMan/utils"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/zmb3/spotify"
)

/* Convert the downloaded audio to MP3 */
func ConvertToMp3(track spotify.FullTrack, playlistName string) string {
	musicDir := "../music/"

	// Check if file already exists
	if utils.CheckFileExistence(track.Name+".mp3", musicDir) {
		fmt.Println("MP3 already exists.")
		return ""
	}

	fmt.Println("Converting to MP3 - ", track.Name)

	outputFile := musicDir + track.Name + ".mp3"

	cmd := exec.Command("ffmpeg", "-i", "../videos/"+track.Name+".mp4", "-vn", "-acodec", "libmp3lame", "-q:a", "2", outputFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error:", err)
		log.Println("ffmpeg output:", string(output))
	}

	fmt.Println("MP3 - ", track.Name+".mp3")
	A_process(track, ".mp3")

	return outputFile
}

/* Add metadata to files */
func AddMetadata(file spotify.FullTrack, extension string, dir string) {
	if ok := utils.CheckFileExistence(file.Name, dir); ok {
		// Check if the file is audio or video
		if extension == ".mp4" {
			V_process(file, extension)
		} else if extension == ".mp3" {
			A_process(file, extension)
		}
	} else {
		fmt.Println("File does not exist")
	}
}

func V_process(track spotify.FullTrack, extension string) {
	videoDir := "../videos"
	mp4 := videoDir + track.Name + extension

	// Add metadata to the mp4 track
	cmd := exec.Command("ffmpeg", "-i", mp4,
		"-metadata", "title="+track.Name,
		"-metadata", "artist="+track.Artists[0].Name,
		"-metadata", "album="+track.Album.Name,
		"-metadata", "track="+strconv.Itoa(track.TrackNumber),
		"-metadata", "year="+strings.Split(track.Album.ReleaseDate, "-")[0],
		"-metadata", "genre="+track.Type,
		"-metadata", "isrc="+track.ExternalIDs["isrc"],
		"-metadata", "disc="+strconv.Itoa(track.DiscNumber),
		"-metadata", "album_artist="+track.Album.Artists[0].Name, "-y", mp4)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Metadata - ", mp4)
}

/* Process audio file */
func A_process(track spotify.FullTrack, extension string) {
	musicDir := "../music/"

	// Download album cover
	imgPath := musicDir + track.Name + ".jpg"
	err := DownloadImage(imgPath, track.Album.Images[0].URL)
	if err != nil {
		fmt.Println("Error downloading image:", err)
		return
	}

	mp3 := musicDir + track.Name + extension

	tempFile := musicDir + track.Name + "_temp" + extension

	// Add metadata to the mp3 track
	cmd := exec.Command("ffmpeg", "-i", imgPath, "-i", mp3,
		"-metadata", "title="+track.Name,
		"-metadata", "artist="+track.Artists[0].Name,
		"-metadata", "album="+track.Album.Name,
		"-metadata", "track="+strconv.Itoa(track.TrackNumber),
		"-metadata", "year="+strings.Split(track.Album.ReleaseDate, "-")[0],
		"-metadata", "isrc="+track.ExternalIDs["isrc"],
		"-metadata", "disc="+strconv.Itoa(track.DiscNumber),
		"-metadata", "album_artist="+track.Album.Artists[0].Name,
		"-map", "0:0", "-map", "1:0", "-c", "copy", "-id3v2_version", "3", "-y", tempFile)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("ffmpeg output:", string(output))
	}

	// Rename temporary output file to the original file name
	if err := os.Rename(tempFile, mp3); err != nil {
		fmt.Println("Error renaming file:", err)
		return
	}

	// Remove the downloaded cover image
	if err := os.Remove(imgPath); err != nil {
		fmt.Println("Error removing cover image:", err)
		return
	}

	fmt.Println("Metadata Applied - ", track.Name+extension)
}
