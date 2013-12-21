package trello

type MemberCreatorData struct {
	Id         string
	AvatarHash string
	FullName   string
	Username   string
}

type UpdateCardData struct {
	ListAfter struct {
		Name string
		Id   string
	}
	ListBefore struct {
		Name string
		Id   string
	}

	Board struct {
		ShortLink string
		Name      string
		Id        string
	}

	Card struct {
		ShortLink string
		IdShort   float64
		Name      string
		Id        string
	}
}

type UpdateCardAction struct {
	Id   string
	Type string
	Date string

	MemberCreator MemberCreatorData
	Data          UpdateCardData
}

type UpdateCardResponse struct {
	Action UpdateCardAction
}
