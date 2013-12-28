package main

import (
	"thcon/trello"
)

func handleCardCreated(body []byte) {
	actionData, err := trello.ParseCreateCardJson(body)
	if err != nil {
		logger.Println("Error encountered while parsing Create Card JSON")
		logger.Println(err.Error())
		return
	}
	hcNotification := &hipchatNotification{notifType: cardCreated, notifData: interface{}(actionData.Action)}
	notifierChannel <- hcNotification
}

func handleCardMoved(body []byte) {
	actionData, err := trello.ParseMoveCardJson(body)
	if err != nil {
		logger.Println("Error encountered while parsing Move Card JSON")
		logger.Println(err.Error())
		return
	}
	hcNotification := &hipchatNotification{notifType: cardMoved, notifData: interface{}(actionData.Action)}
	notifierChannel <- hcNotification
}

func handleListOpened(body []byte) {
	actionData, err := trello.ParseOpenedListJson(body)
	if err != nil {
		logger.Println("Error encountered while parsing Opened List JSON")
		logger.Println(err.Error())
		return
	}
	hcNotification := &hipchatNotification{notifType: listOpened, notifData: interface{}(actionData.Action)}
	notifierChannel <- hcNotification
}
func handleListClosed(body []byte) {
	actionData, err := trello.ParseClosedListJson(body)
	if err != nil {
		logger.Println("Error encountered while parsing Closed List JSON")
		logger.Println(err.Error())
		return
	}
	hcNotification := &hipchatNotification{notifType: listClosed, notifData: interface{}(actionData.Action)}
	notifierChannel <- hcNotification
}

func handleListCreated(body []byte) {
	actionData, err := trello.ParseCreateListJson(body)
	if err != nil {
		logger.Println("Error encountered while parsing Create List JSON")
		logger.Println(err.Error())
		return
	}
	hcNotification := &hipchatNotification{notifType: listCreated, notifData: interface{}(actionData.Action)}
	notifierChannel <- hcNotification
}

func handleMemberAddedToCard(body []byte) {
	actionData, err := trello.ParseAddMemberToCardJson(body)
	if err != nil {
		logger.Println("Error encountered while parsing Add Member to Card JSON")
		logger.Println(err.Error())
		return
	}
	hcNotification := &hipchatNotification{notifType: memberAddedToCard, notifData: interface{}(actionData.Action)}
	notifierChannel <- hcNotification
}
