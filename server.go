package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"thcon/trello"
	"time"
)

func startServer(config *thConfigData) error {
	handler := http.NewServeMux()
	server := &http.Server{
		Addr:           config.App.ServerAddress,
		Handler:        handler,
		ReadTimeout:    time.Duration(config.App.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.App.ReadTimeout) * time.Second,
		MaxHeaderBytes: 0,
	}

	// setup handlers
	handler.HandleFunc("/", rootHandler)
	handler.HandleFunc("/board", boardUpdateHandler)

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func rootHandler(respWriter http.ResponseWriter, request *http.Request) {
	respWriter.Write([]byte("TODO"))
}

func boardUpdateHandler(respWriter http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	request.Body.Close()

	actionData, _ := trello.ParseUpdateCardJson(reqBody)
	if actionData.Action.Type == "updateCard" {
		fmt.Printf("%#v\n", actionData)
	}

	// always reply with 200 to Trello. Any errors are to be logged
	respWriter.WriteHeader(200)
	respWriter.Write([]byte{})
}
