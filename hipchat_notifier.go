package main

import (
	"fmt"
	"github.com/hoisie/mustache"
	"thcon/hipchat"
	"thcon/trello"
	"time"
)

func newHipchatNotifier(room *hipchat.Room, sleepDuration int) *hipchatNotifier {
	hn := &hipchatNotifier{}
	hn.inputChan = make(chan *hipchatNotification, 100)
	hn.quitChan = make(chan int)
	hn.room = room
	hn.sleepDuration = sleepDuration

	return hn
}

func (hn *hipchatNotifier) startHipchatNotifier() {
NotifierLoop:
	for {
		select {
		case notif := <-hn.inputChan:
			allNotifications := make([]*hipchatNotification, 0)
			allNotifications = append(allNotifications, notif)
			for len(hn.inputChan) > 0 {
				notif = <-hn.inputChan
				allNotifications = append(allNotifications, notif)
			}
			err := hn.sendNotifications(allNotifications)
			if err != nil {
				logger.Println("Unable to send notifications to hipchat")
				logger.Println(err.Error())
			}
		case <-hn.quitChan:
			logger.Println("Received quit signal on quit channel. Quiting")
			break NotifierLoop
		default:
			time.Sleep(time.Duration(hn.sleepDuration) * time.Second)
		}
	}
}

func (hn *hipchatNotifier) sendNotifications(notifications []*hipchatNotification) error {
	// group actions by member ids
	memberActions := make(map[string][]string)
	memberNames := make(map[string]string)
	for _, notif := range notifications {
		memberData, actionString := notif.actionToString()
		if _, ok := memberNames[memberData.Id]; !ok {
			memberNames[memberData.Id] = memberData.FullName
		}
		if _, ok := memberActions[memberData.Id]; !ok {
			memberActions[memberData.Id] = make([]string, 0)
		}
		memberActions[memberData.Id] = append(memberActions[memberData.Id], actionString)
	}

	type actionContext struct {
		ActionString string
	}
	type userContext struct {
		Username string
		Actions  []actionContext
	}
	type templateContext struct {
		BoardName string
		Users     []userContext
	}

	// prepare sturcts for mustache templtes
	userContexts := make([]userContext, 0)
	for memberId, memberName := range memberNames {
		uc := userContext{Username: memberName, Actions: make([]actionContext, 0)}
		for _, action := range memberActions[memberId] {
			ac := actionContext{ActionString: action}
			uc.Actions = append(uc.Actions, ac)
		}
		userContexts = append(userContexts, uc)
	}
	context := templateContext{BoardName: hn.room.RoomName, Users: userContexts}
	notificationString := mustache.Render(notificationHtml, context)

	if err := hn.room.PostMessage(notificationString, dFrom, dFormat, dNotify,
		dColor); err != nil {
		return err
	}
	return nil
}

func (notif *hipchatNotification) actionToString() (trello.MemberCreatorData, string) {
	var actionString string
	var memberData trello.MemberCreatorData

	switch notif.notifType {
	case cardCreated:
		action := notif.notifData.(trello.CreateCardAction)
		cardName := action.Data.Card.Name
		listName := action.Data.List.Name
		actionString = fmt.Sprintf(cardCreatedMsg, cardName, listName)
		memberData = action.MemberCreator
	case cardMoved:
		action := notif.notifData.(trello.MoveCardAction)
		cardName := action.Data.Card.Name
		listBeforeName := action.Data.ListBefore.Name
		listAfterName := action.Data.ListAfter.Name
		actionString = fmt.Sprintf(cardMovedMsg, cardName, listBeforeName,
			listAfterName)
		memberData = action.MemberCreator
	case listOpened:
		action := notif.notifData.(trello.OpenedListAction)
		listName := action.Data.List.Name
		actionString = fmt.Sprintf(listOpenedMsg, listName)
		memberData = action.MemberCreator
	case listClosed:
		action := notif.notifData.(trello.ClosedListAction)
		listName := action.Data.List.Name
		actionString = fmt.Sprintf(listClosedMsg, listName)
		memberData = action.MemberCreator
	case listCreated:
		action := notif.notifData.(trello.CreateListAction)
		listName := action.Data.List.Name
		actionString = fmt.Sprintf(listCreatedMsg, listName)
		memberData = action.MemberCreator
	case memberAddedToCard:
		action := notif.notifData.(trello.AddMemberToCardAction)
		cardName := action.Data.Card.Name
		memberName := action.Member.FullName
		actionString = fmt.Sprintf(memberAddedToCardMsg, memberName, cardName)
		memberData = action.MemberCreator
	}

	return memberData, actionString
}
