package trello

import (
	"time"
)

type MemberCreatorData struct {
	Id         string
	AvatarHash string
	FullName   string
	Username   string
}

type Board struct {
	ShortLink string
	Name      string
	Id        string
}

type Card struct {
	ShortLink string
	IdShort   float64
	Name      string
	Id        string
}

type List struct {
	Name string
	Id   string
}
type OpenedList struct {
	Name   string
	Id     string
	Closed struct {
		Id     string `json:"_id"`
		Name   string
		Closed bool
	}
}
type ClosedList struct {
	Name   string
	Id     string
	Closed bool
}

type CreateCardData struct {
	Board `json:"board"`
	List  `json:"list"`
	Card  `json:"card"`
}
type CreateCardAction struct {
	Id            string
	Type          string
	Date          time.Time
	MemberCreator MemberCreatorData

	Data CreateCardData
}
type CreateCardResponse struct {
	Action CreateCardAction
}

type MoveCardData struct {
	ListAfter  List
	ListBefore List
	Board      `json:"board"`
	Card       `json:"card"`
}
type MoveCardAction struct {
	Id            string
	Type          string
	Date          time.Time
	MemberCreator MemberCreatorData

	Data MoveCardData
}
type MoveCardResponse struct {
	Action MoveCardAction
}

type OpenedListData struct {
	Board `json:"board"`
	List  OpenedList
	Old   struct {
		Closed bool
	}
}
type ClosedListData struct {
	Board `json:"board"`
	List  ClosedList
	Old   struct {
		Closed bool
	}
}
type OpenedListAction struct {
	Id            string
	Type          string
	Date          time.Time
	MemberCreator MemberCreatorData

	Data OpenedListData
}
type ClosedListAction struct {
	Id            string
	Type          string
	Date          time.Time
	MemberCreator MemberCreatorData

	Data ClosedListData
}
type OpenedListResponse struct {
	Action OpenedListAction
}
type ClosedListResponse struct {
	Action ClosedListAction
}

type CreateListData struct {
	Board `json:"board"`
	List  `json:"list"`
}
type CreateListAction struct {
	Id            string
	Type          string
	Date          time.Time
	MemberCreator MemberCreatorData

	Data CreateListData
}
type CreateListResponse struct {
	Action CreateListAction
}

type AddMemberToCardData struct {
	Board    `json:"board"`
	Card     `json:"card"`
	IdMember string
}
type AddMemberToCardAction struct {
	Id            string
	Type          string
	Date          time.Time
	MemberCreator MemberCreatorData

	Member MemberCreatorData
	Data   AddMemberToCardData
}
type AddMemberToCardResponse struct {
	Action AddMemberToCardAction
}
