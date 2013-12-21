package main

import (
	"fmt"
	"os"

	_ "thcon/hipchat"
	"thcon/trello"
)

func main() {
	config := getConfig()

	fmt.Printf("Starting http server at: %s\n", config.App.ServerAddress)
	startServer(config)
	fmt.Print("Server finished\n")
}

func getConfig() *thConfigData {
	config, err := readConfig(configFileName)
	if err != nil {
		fmt.Printf("Error while reading config file.\n%s\n", err.Error())
		os.Exit(1)
	}
	if len(config.Hipchat.ApiKey) == 0 || len(config.Hipchat.RoomId) == 0 ||
		len(config.Trello.ApiKey) == 0 || len(config.Trello.ApiSecret) == 0 {
		fmt.Printf("Please enter details for the Trello and Hipchat Api access in your config file.")
		os.Exit(1)
	}
	if len(config.Trello.AuthToken) == 0 {
		fmt.Printf("Enter details for your Trello token by visiting: %s\n",
			trello.GetAuthUrl(config.Trello.ApiKey))
		os.Exit(1)
	}
	return config
}
