package user

import (
	"context"

	"SongMan/utils"

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
