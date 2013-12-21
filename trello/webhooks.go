package trello

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type webhook struct {
	Id          string
	Description string
	IdModel     string
	CallbackURL string
	Active      bool
}

func (api *TrelloApi) GetWebhooks() ([]webhook, error) {
	getWebhooksUrl := createCompleteUrlFromPath(webhooksListPath, api.authToken)
	getWebhooksUrl, err := api.addAuthKeysToStringUrl(getWebhooksUrl)
	if err != nil {
		return nil, err
	}

	resp, body, err := getAndReadResponse(getWebhooksUrl)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, createRequestFailedError(resp, body)
	}

	webhooks := make([]webhook, 5)
	err = json.Unmarshal(body, &webhooks)
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

func (api *TrelloApi) CreateWebhook(webhook *Webhook) error {
	jsonWebhook, err := json.Marshal(webhook)
	if err != nil {
		return err
	}

	createWebhookUrl := createCompleteUrlFromPath(webhookCreatePath)
	createWebhookUrl, err = api.addAuthKeysToStringUrl(createWebhookUrl)
	if err != nil {
		return err
	}

	requestBody := bytes.NewReader(jsonWebhook)
	resp, body, err := postAndReadResponse(createWebhookUrl, "application/json", requestBody)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		// if this webhook already exists, Trello replies with 400
		if resp.StatusCode == http.StatusBadRequest {
			if strings.TrimSpace(string(body)) == webhookExistsErrorMessage {
				return nil
			}
		}
		return createRequestFailedError(resp, body)
	}

	err = json.Unmarshal(body, webhook)
	if err != nil {
		return err
	}
	return nil
}
