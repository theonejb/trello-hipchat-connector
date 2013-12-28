package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"thcon/hipchat"
	"thcon/trello"
)

var logger *log.Logger
var notifierChannel chan *hipchatNotification

func main() {
	logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	config := getConfig()

	// setup channels
	serverStopSignal := make(chan int)
	ctrlCSignal := make(chan os.Signal)

	// setup control-c signal and goroutine
	signal.Notify(ctrlCSignal, syscall.SIGINT)
	go stopServer(ctrlCSignal, serverStopSignal)

	// setup hipchat notifier
	hipchatRoom := hipchat.NewRoom(config.Hipchat.ApiKey, config.Hipchat.RoomId,
		config.Hipchat.RoomName, logger)
	hipchatNotifier := newHipchatNotifier(hipchatRoom, config.App.SleepDuration)
	notifierChannel = hipchatNotifier.inputChan
	go hipchatNotifier.startHipchatNotifier()

	logger.Println("Starting HTTP server")
	if err := startServer(config, serverStopSignal); err != nil {
		logger.Println("startServer returned an error")
		logger.Println(err.Error())
	}
	logger.Println("HTTP server stopped")

	logger.Println("Stopping hipchat notifier. Might take upto a minute.")
	hipchatNotifier.quitChan <- 1
}

func stopServer(stopSignal <-chan os.Signal, serverStop chan<- int) {
	<-stopSignal
	serverStop <- 1
}

func getConfig() *thConfigData {
	config, err := readConfig(configFileName)
	if err != nil {
		logger.Fatalf("Error while reading config file.\n%s\n", err.Error())
	}
	if len(config.Hipchat.ApiKey) == 0 || len(config.Hipchat.RoomId) == 0 ||
		(len(config.Hipchat.RoomName) == 0) || len(config.Trello.ApiKey) == 0 ||
		len(config.Trello.ApiSecret) == 0 {
		logger.Fatalf("Please enter details for the Trello and Hipchat Api access in your config file.")
	}
	if len(config.Trello.AuthToken) == 0 {
		logger.Fatalf("Enter details for your Trello token by visiting: %s\n", trello.GetAuthUrl(config.Trello.ApiKey))
	}
	return config
}
