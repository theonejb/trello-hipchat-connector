package trello

import (
	"encoding/json"
)

func ParseUpdateCardJson(response []byte) (*UpdateCardResponse, error) {
	actionResponse := &UpdateCardResponse{}
	err := json.Unmarshal(response, actionResponse)
	if err != nil {
		return nil, err
	}

	return actionResponse, nil
}
