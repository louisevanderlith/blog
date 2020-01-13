package droxo

import "golang.org/x/oauth2"

/*
var (
	Config = oauth2.Config{
		ClientID:     "",
		ClientSecret: "blogsecret",
		Scopes:       []string{"blog", "comment"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://oauth2:8086/auth",
			TokenURL: "http://oauth2:8086/token",
		},
	}
)
*/
var AuthServer = oauth2.Endpoint{
	AuthURL:   "http://oauth2:8086/auth",
	TokenURL:  "http://oauth2:8086/token",
	AuthStyle: 0,
}