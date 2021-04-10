package google

import "golang.org/x/oauth2"

type GoogleOauth2 interface {
	GetConnect() *oauth2.Config
}
