package main

import (
	"thcon/hipchat"
)

type hipchatNotifier struct {
	inputChan     chan *hipchatNotification
	quitChan      chan int
	room          *hipchat.Room
	sleepDuration int
}

type hipchatNotification struct {
	notifType int
	notifData interface{}
}
