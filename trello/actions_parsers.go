package trello

import (
	"encoding/json"
)

func ParseCreateCardJson(response []byte) (*CreateCardResponse, error) {
	actionResponse := &CreateCardResponse{}
	err := json.Unmarshal(response, actionResponse)
	if err != nil {
		return nil, err
	}
	return actionResponse, nil
}

func ParseMoveCardJson(response []byte) (*MoveCardResponse, error) {
	actionResponse := &MoveCardResponse{}
	err := json.Unmarshal(response, actionResponse)
	if err != nil {
		return nil, err
	}
	return actionResponse, nil
}

func ParseOpenedListJson(response []byte) (*OpenedListResponse, error) {
	actionResponse := &OpenedListResponse{}
	err := json.Unmarshal(response, actionResponse)
	if err != nil {
		return nil, err
	}
	return actionResponse, nil
}
func ParseClosedListJson(response []byte) (*ClosedListResponse, error) {
	actionResponse := &ClosedListResponse{}
	err := json.Unmarshal(response, actionResponse)
	if err != nil {
		return nil, err
	}
	return actionResponse, nil
}

func ParseCreateListJson(response []byte) (*CreateListResponse, error) {
	actionResponse := &CreateListResponse{}
	err := json.Unmarshal(response, actionResponse)
	if err != nil {
		return nil, err
	}
	return actionResponse, nil
}

func ParseAddMemberToCardJson(response []byte) (*AddMemberToCardResponse, error) {
	actionResponse := &AddMemberToCardResponse{}
	err := json.Unmarshal(response, actionResponse)
	if err != nil {
		return nil, err
	}
	return actionResponse, nil
}
