package bp

import (
	"github.com/zmb3/spotify"
)

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
