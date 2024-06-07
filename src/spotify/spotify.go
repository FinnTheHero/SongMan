package spotify

import (
	"SongMan/utils"
	"context"
	"regexp"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

/* Retrieve configuration credentials from the environment variables. */
func GetConfigCredentials() *clientcredentials.Config {
	// Configure client credentials
	authConfig := &clientcredentials.Config{
		ClientID:     utils.GetEnvVariables("CLIENT_ID"),
		ClientSecret: utils.GetEnvVariables("CLIENT_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	return authConfig
}

/* Create a new user client using configuration credentials. */
func GetNewUser(authConfig *clientcredentials.Config) spotify.Client {
	// Get the access token
	accessToken, err := authConfig.Token(context.Background())

	if err != nil {
		panic(err)
	}

	// Create a new client
	client := spotify.Authenticator{}.NewClient(accessToken)

	return client

}

func HandleSpotify() spotify.Client {
	// Configure client credentials
	authConfig := GetConfigCredentials()

	// Create a new client
	client := GetNewUser(authConfig)

	return client
}

/* Return full playlist details using client and playlist ID. */
func GetPlaylist(client spotify.Client, playlistID spotify.ID) *spotify.FullPlaylist {
	// Get the playlist ID
	p := spotify.ID(playlistID)

	// Get the playlist
	playlist, err := client.GetPlaylist(p)

	if err != nil {
		panic(err)
	}

	return playlist
}

/* Return the list of tracks from the playlist. */
func GetTracks(playlist *spotify.FullPlaylist) spotify.PlaylistTrackPage {
	// Get the playlist tracks
	tracks := playlist.Tracks

	return tracks
}

/* Get single track */
func GetTrack(client spotify.Client, trackID spotify.ID) spotify.FullTrack {
	// Get the track ID
	t := spotify.ID(trackID)

	// Get the track
	track, err := client.GetTrack(t)
	if err != nil {
		panic(err)
	}

	return *track
}

/* Extracts the Spotify track or playlist ID and mode from a given link */
func ExtractSpotifyID(link string) (spotify.ID, string) {
	// Regular expression patterns for track and playlist links
	trackPattern := regexp.MustCompile(`^https?:\/\/open.spotify.com\/track\/([a-zA-Z0-9]+)`)
	playlistPattern := regexp.MustCompile(`^https?:\/\/open.spotify.com\/playlist\/([a-zA-Z0-9]+)`)

	if trackPattern.MatchString(link) {
		// Extract the track ID from the link
		matches := trackPattern.FindStringSubmatch(link)
		return spotify.ID(matches[1]), "track"
	} else if playlistPattern.MatchString(link) {
		// Extract the playlist ID from the link
		matches := playlistPattern.FindStringSubmatch(link)
		return spotify.ID(matches[1]), "playlist"
	} else {
		// Return an empty ID if the link is not a valid track or playlist link
		panic("Invalid link")
	}
}
