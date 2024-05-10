package main

func main() {

	// Configure client credentials
	authConfig := GetConfigCredentials()

	// Create a new client
	client := GetNewUser(authConfig)

	blueprint := GeneratePlaylistBlueprint(client, "3cEYpjA9oz9GiPac4AsH4n")

	ExportBlueprint(blueprint)
}
