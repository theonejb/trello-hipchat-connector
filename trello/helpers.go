package trello

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func createRequestFailedError(resp *http.Response, body []byte) error {
	return errors.New(fmt.Sprintf("Trello api request failed. Code: %d Body: %s",
		resp.StatusCode, body))
}

func getAndReadResponse(urlString string) (*http.Response, []byte, error) {
	resp, err := http.Get(urlString)
	if err != nil {
		return nil, nil, err
	}
	body := make([]byte, resp.ContentLength)
	if _, err = resp.Body.Read(body); err != nil {
		return nil, nil, err
	}
	resp.Body.Close()
	return resp, body, nil
}

func postAndReadResponse(urlString string, bodyType string, body io.Reader) (*http.Response, []byte, error) {
	resp, err := http.Post(urlString, bodyType, body)
	if err != nil {
		return nil, nil, err
	}
	respBody := make([]byte, resp.ContentLength)
	if _, err = resp.Body.Read(respBody); err != nil {
		return nil, nil, err
	}
	resp.Body.Close()
	return resp, respBody, nil
}

func createCompleteUrlFromPath(path string, params ...interface{}) string {
	var completePath string
	if len(params) != 0 {
		completePath = fmt.Sprintf(path, params...)
	} else {
		completePath = path
	}
	completeUrl := apiBaseUrl + completePath
	return completeUrl
}
