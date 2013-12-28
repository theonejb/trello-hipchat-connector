package main

const configFileName string = "Programming/thConfig.json"

// default values for the message options
const dFrom string = "TrelloConnector"
const dFormat string = "html"
const dNotify string = "1"
const dColor string = "yellow"

// response types from the Trello API webhooks
const (
	cardCreated       = iota
	cardMoved         = iota
	listOpened        = iota
	listClosed        = iota
	listCreated       = iota
	memberAddedToCard = iota
	unknownResponse   = iota
)

// notification message strings
const cardCreatedMsg string = "Created card [%s] in list [%s]"
const cardMovedMsg string = "Moved card [%s] from list [%s] to [%s]"
const listOpenedMsg string = "Unarchived list [%s]"
const listClosedMsg string = "Archived list [%s]"
const listCreatedMsg string = "Created a new list [%s]"
const memberAddedToCardMsg string = "Added [%s] to card [%s]"

// notifications template
const notificationHtml string = `
Updates to board {{BoardName}}:
	<ul>
	{{#Users}}
	<li>
	By member {{Username}}:
	<ul>
	{{#Actions}}
	<li>{{ActionString}}</li>
	{{/Actions}}
	</ul>
	</li>
	{{/Users}}
	</ul>
`
