package gotifyservice

import (
	"net/http"
	"net/url"

	"github.com/gotify/go-api-client/v2/auth"
	"github.com/gotify/go-api-client/v2/client/message"
	"github.com/gotify/go-api-client/v2/gotify"
	"github.com/gotify/go-api-client/v2/models"
)

type gotifyService struct {
	gotifyURL        string
	applicationToken string
}

type GotifyService interface {
	SendMessage(msg *models.MessageExternal) error
}

func (gs *gotifyService) SendMessage(msg *models.MessageExternal) error {
	myURL, _ := url.Parse(gs.gotifyURL)
	client := gotify.NewClient(myURL, &http.Client{})
	_, err := client.Version.GetVersion(nil)
	if err != nil {
		return err
	}

	params := message.NewCreateMessageParams()
	params.Body = msg
	_, err = client.Message.CreateMessage(params, auth.TokenAuth(gs.applicationToken))

	if err != nil {
		return err
	}
	return nil
}

func NewGotifyService(gotifyURL, applicationToken string) GotifyService {
	return &gotifyService{
		gotifyURL:        gotifyURL,
		applicationToken: applicationToken,
	}
}
