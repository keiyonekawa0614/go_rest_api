package controllers

import (
	"context"
	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
)

type GoogleOauth2Controller struct {
	Config *oauth2.Config
}

func NewGoogleOauth2Controller(config *oauth2.Config) *GoogleOauth2Controller {
	return &GoogleOauth2Controller {
		Config: config,
	}
}

// google oauth2 redirect.
func (controller *GoogleOauth2Controller) Auth(c Context) (err error) {
	url := controller.Config.AuthCodeURL("")
	return c.Redirect(302, url)
}

// google call back method.
func (controller *GoogleOauth2Controller) Callback(c Context) (err error) {
	context := context.Background()
	tok, err := controller.Config.Exchange(context, c.QueryParam("code"))
	if err != nil {
		return c.JSON(500, NewError(err))
	}
	if tok.Valid() == false {
		return c.JSON(500, NewError(err))
	}
	service, _ := v2.New(controller.Config.Client(context, tok))
	tokenInfo, _ := service.Tokeninfo().AccessToken(tok.AccessToken).Context(context).Do()

	var object = make(map[string]string)
	object["access_token"] = tok.AccessToken
	object["user_id"] = tokenInfo.UserId
	object["email"] = tokenInfo.Email

	return c.JSON(200, object)
}
