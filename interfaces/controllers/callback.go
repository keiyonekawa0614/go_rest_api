package controllers

import (
	"context"
	"app/infrastructure/google"
	v2 "google.golang.org/api/oauth2/v2"
)

type CallbackController struct {
}

func NewCallbackController() * CallbackController {
	return &CallbackController{}
}

func (controller *CallbackController) Get(c Context) (err error) {
	config := google.GetConnect()
	context := context.Background()
	tok, err := config.Exchange(context, c.QueryParam("code"))
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if tok.Valid() == false {
		c.JSON(500, NewError(err))
		return
	}
	service, _ := v2.New(config.Client(context, tok))
	tokenInfo, _ := service.Tokeninfo().AccessToken(tok.AccessToken).Context(context).Do()

	var object = make(map[string]string)
	object["user_id"] = tokenInfo.UserId
	object["email"] = tokenInfo.Email

	return c.JSON(200, object)
}
