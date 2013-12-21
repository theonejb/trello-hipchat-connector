package trello

import (
	"errors"
	"fmt"
	"net/url"
)

func NewApi(key, token string) (*TrelloApi, error) {
	api := &TrelloApi{apiKey: key, authToken: token}
	return api, nil
}

func (api *TrelloApi) addAuthKeysToStringUrl(oldUrl string) (string, error) {
	parsedUrl, err := url.Parse(oldUrl)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to add auth keys to URL. URL: %s", oldUrl))
	}

	query := parsedUrl.Query()
	query.Add("key", api.apiKey)
	query.Add("token", api.authToken)

	parsedUrl.RawQuery = query.Encode()
	return parsedUrl.String(), nil
}

func GetAuthUrl(apiKey string) string {
	urlParams := url.Values{"scope": {"read,write"}, "expiration": {authTokenExpiration},
		"name": {applicationName}, "key": {apiKey}, "response_type": {"token"}}
	return fmt.Sprintf("%s?%s", authTokenGetUrl, urlParams.Encode())
}
