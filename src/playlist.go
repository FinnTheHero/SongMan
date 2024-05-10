package main

import "github.com/zmb3/spotify"

/*
Return full playlist details using client and playlist ID.
*/
func GetPlaylist(client spotify.Client, playlistIDString spotify.ID) *spotify.FullPlaylist {
	// Get the playlist ID
	playlistID := spotify.ID(playlistIDString)

	// Get the playlist
	playlist, err := client.GetPlaylist(playlistID)

	if err != nil {
		panic(err)
	}

	return playlist
}
