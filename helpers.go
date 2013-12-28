package main

import (
	"encoding/json"
)

func getWebhookCallbackType(body []byte) int {
	decodedJson := make(map[string]interface{})
	err := json.Unmarshal(body, &decodedJson)
	if err != nil {
		return unknownResponse
	}

	if val, ok := decodedJson["action"]; ok {
		if actionType, ok := val.(map[string]interface{})["type"]; ok {
			if actionDataDict, ok := getActionDataDict(val); ok {
				switch actionType {
				case "createCard":
					return getCreateCardActionType(actionDataDict)
				case "updateCard":
					return getUpdateCardActionType(actionDataDict)
				case "updateList":
					return getUpdatelistActionType(actionDataDict)
				case "createList":
					return getListCreatedActionType(actionDataDict)
				case "addMemberToCard":
					return getAddMemberToCardActionType(actionDataDict)
				default:
					logger.Printf("Unhandled action type: %s\nAction Data:\n", actionType)
					logger.Printf("%#v\n", actionDataDict)
				}
			}
		}
	}

	return unknownResponse
}

func getActionDataDict(actionDict interface{}) (map[string]interface{}, bool) {
	if actionData, ok := actionDict.(map[string]interface{})["data"]; ok {
		actionDataDict := actionData.(map[string]interface{})
		return actionDataDict, true
	}
	return nil, false
}

func getCreateCardActionType(actionDataDict map[string]interface{}) int {
	return cardCreated
}

func getUpdateCardActionType(actionDataDict map[string]interface{}) int {
	_, listAfterExists := actionDataDict["listAfter"]
	_, listBeforeExists := actionDataDict["listBefore"]
	if listAfterExists && listBeforeExists {
		return cardMoved
	}

	logger.Println("Unable to guess updateCard sub-action type. Action Data:")
	logger.Printf("%#v", actionDataDict)
	return unknownResponse
}

func getUpdatelistActionType(actionDataDict map[string]interface{}) int {
	_, listExists := actionDataDict["list"]
	_, oldStateExists := actionDataDict["old"]
	wasListClosed := !actionDataDict["old"].(map[string]interface{})["closed"].(bool)

	if listExists && oldStateExists {
		if wasListClosed {
			return listClosed
		} else {
			return listOpened
		}
	}

	logger.Println("Unable to guess updateList sub-action type. Action Data:")
	logger.Printf("%#v", actionDataDict)
	return unknownResponse
}

func getListCreatedActionType(actionDataDict map[string]interface{}) int {
	return listCreated
}

func getAddMemberToCardActionType(actionDataDict map[string]interface{}) int {
	return memberAddedToCard
}
