package trello

const applicationName = "Trello Hipchat Connector"

const authTokenGetUrl string = "https://trello.com/1/authorize"
const authTokenExpiration string = "never"

const apiBaseUrl string = "https://trello.com/1/"
const webhooksListPath string = "tokens/%s/webhooks"
const webhookCreatePath string = "webhooks"

const webhookExistsErrorMessage string = "A webhook with that callback, model, and token already exists"
