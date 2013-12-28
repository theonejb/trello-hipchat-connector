package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func startServer(config *thConfigData, exitChan <-chan int) error {
	lAddress, err := net.ResolveTCPAddr("tcp", config.App.ServerAddress)
	if err != nil {
		logger.Println("Encountered error while resolving server address")
		return err
	}
	listener, err := net.ListenTCP("tcp", lAddress)
	listenerClosed := false
	if err != nil {
		logger.Println("Encountered error while creating TCP listener")
		return err
	}

	handler := http.NewServeMux()
	server := &http.Server{
		Handler:        handler,
		ReadTimeout:    time.Duration(config.App.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.App.ReadTimeout) * time.Second,
		MaxHeaderBytes: 0,
	}

	// setup handlers
	handler.HandleFunc("/", rootHandler)
	handler.HandleFunc("/board", boardUpdateHandler)

	// setup goroutine to listen for exit command
	go func() {
		<-exitChan
		listener.Close()
		listenerClosed = true
		logger.Println("Listener stopped. Server will now exit")
	}()

	err = server.Serve(listener)
	// if we closed the listener ourselves, no need to report an error
	if err != nil && !listenerClosed {
		logger.Println("Server.Serve returned an error")
		return err
	}
	return nil
}

func rootHandler(respWriter http.ResponseWriter, request *http.Request) {
	respWriter.Write([]byte("TODO"))
}

func boardUpdateHandler(respWriter http.ResponseWriter, request *http.Request) {
	reqBody, err := ioutil.ReadAll(request.Body)
	request.Body.Close()
	if err != nil {
		logger.Println("Error encountered while reading data from request")
		logger.Println(err.Error())
		return
	}

	switch getWebhookCallbackType(reqBody) {
	case cardCreated:
		handleCardCreated(reqBody)
	case cardMoved:
		handleCardMoved(reqBody)
	case listOpened:
		handleListOpened(reqBody)
	case listClosed:
		handleListClosed(reqBody)
	case listCreated:
		handleListCreated(reqBody)
	case memberAddedToCard:
		handleMemberAddedToCard(reqBody)
	case unknownResponse:
		logger.Println("Unknown action")
	}

	// always reply with 200 to Trello. Any errors are to be logged
	respWriter.WriteHeader(200)
	respWriter.Write([]byte{})
}
