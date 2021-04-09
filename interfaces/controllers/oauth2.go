package controllers

import "app/infrastructure/google"

type Oauth2Controller struct {
}

func NewOauth2Controller() *Oauth2Controller {
	return &Oauth2Controller{}
}

func (controller *Oauth2Controller) Get(c Context) (err error) {
	config := google.GetConnect()
	url := config.AuthCodeURL("")
	return c.Redirect(302, url)
}
