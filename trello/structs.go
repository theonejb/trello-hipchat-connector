package trello

type TrelloApi struct {
	apiKey    string
	authToken string
}

type Webhook struct {
	Id          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	CallbackURL string `json:"callbackURL"`
	IdModel     string `json:"idModel"`
	Active      bool   `json:"-"`
}
