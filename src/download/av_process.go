package download

import (
	"SongMan/utils"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/zmb3/spotify"
)

/* Convert the downloaded audio to MP3 */
func ConvertToMp3(inputFile string) string {
	// Create music file direcotry if it doesn't exist
	musicDir := "../music/"

	err := os.MkdirAll(musicDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	fmt.Println("Converting to MP3 - ", inputFile)

	outputFile := musicDir + inputFile + ".mp3"

	cmd := exec.Command("ffmpeg", "-i", "../videos/"+inputFile+".mp4", "-vn", "-acodec", "libmp3lame", "-q:a", "2", outputFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("ffmpeg output:", string(output))
	}

	fmt.Println("MP3 - ", inputFile+".mp3")
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

func V_process(file spotify.FullTrack, extension string) {
	// Add metadata to the mp4 file
	videoDir := "../videos"
	mp4 := videoDir + file.Name + extension

	// cmd := exec.Command("ffmpeg", "-i", mp4, "-metadata", "title=YourTitle", "-metadata", "artist=YourArtist", "-y", mp4)
	cmd := exec.Command("ffmpeg", "-i", mp4, "-metadata", "title="+file.Name, "-metadata", "artist="+file.Artists[0].Name, "-metadata", "album="+file.Album.Name, "-metadata", "track="+strconv.Itoa(file.TrackNumber), "-metadata", "year="+strings.Split(file.Album.ReleaseDate, "-")[0], "-metadata", "genre="+file.Type, "-metadata", "isrc="+file.ExternalIDs["isrc"], "-metadata", "disc="+strconv.Itoa(file.DiscNumber), "-metadata", "-metadata", "album_artist="+file.Album.Artists[0].Name, "-y", mp4)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Metadata - ", mp4)
}

/* Process audio file */
func A_process(file spotify.FullTrack, extension string) {
	// Add metadata to the mp3 file
	musicDir := "../music/"

	// Download album cover
	imgPath := musicDir + file.Name + ".jpg"
	err := utils.DownloadFile(imgPath, file.Album.Images[0].URL)
	if err != nil {
		fmt.Println("Error downloading image:", err)
		return
	}

	mp3 := musicDir + file.Name + extension

	tempFile := musicDir + file.Name + "_temp" + extension

	cmd := exec.Command("ffmpeg", "-i", imgPath, "-i", mp3,
		"-metadata", "title="+file.Name,
		"-metadata", "artist="+file.Artists[0].Name,
		"-metadata", "album="+file.Album.Name,
		"-metadata", "track="+strconv.Itoa(file.TrackNumber),
		"-metadata", "year="+strings.Split(file.Album.ReleaseDate, "-")[0],
		"-metadata", "isrc="+file.ExternalIDs["isrc"],
		"-metadata", "disc="+strconv.Itoa(file.DiscNumber),
		"-metadata", "album_artist="+file.Album.Artists[0].Name,
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

	fmt.Println("Metadata - ", file.Name+extension)
}
