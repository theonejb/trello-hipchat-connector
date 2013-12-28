package main

import (
	"encoding/json"
	"os"
	"os/user"
	"path"
)

type trelloBoardConfig struct {
	Id string
}
type trelloConfig struct {
	ApiKey    string
	ApiSecret string
	AuthToken string
	Boards    trelloBoardConfig
}
type hipchatConfig struct {
	ApiKey   string
	RoomId   string
	RoomName string
}
type appConfig struct {
	BoardWebhookCallbackUrl string
	IdModel                 string
	ServerAddress           string
	ReadTimeout             int
	WriteTimeout            int
	SleepDuration           int
}
type thConfigData struct {
	Trello  trelloConfig
	Hipchat hipchatConfig
	App     appConfig
}

func readConfig(fileName string) (*thConfigData, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	homeDir := usr.HomeDir
	fullFilePath := path.Join(homeDir, fileName)
	configFile, err := os.OpenFile(fullFilePath, os.O_RDONLY, 0)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}

	jsonDecoder := json.NewDecoder(configFile)
	config := &thConfigData{}
	if err := jsonDecoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
