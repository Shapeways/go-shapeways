package shapeways_test

import (
	"github.com/Shapeways/go-shapeways/shapeways"
)

func ExampleNewClient() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)

	// If the user has been authed before then
	// explicitly set their token/secret
	client.OauthCredentials.Token = "USERS TOKEN"
	client.OauthCredentials.Secret = "USERS SECRET"
}
