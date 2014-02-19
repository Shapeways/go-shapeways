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

func ExampleClient_Connect() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)

	authUrl, err := client.Connect()
	if err != nil {
		// handle error
	}

	// access oauth secret
	client.OauthCredentials.Secret

	// send user to authUrl
}

func ExampleClient_Verify() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)

	// parse token and verifier from oauth callback
	err := client.Verify(token, verifier)
	if err != nil {
		// handle error
	}

	// access oauth token/secret
	client.OauthCredentials.Token
	client.OauthCredentials.Secret
}

func ExampleClient_VerifyURL() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)

	// raw callback url from oauth callback
	err := client.VerifyURL(callbackUrl)
	if err != nil {
		// handle error
	}

	// access oauth token/secret
	client.OauthCredentials.Token
	client.OauthCredentials.Secret
}

func ExampleClient_Url() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	fmt.Printf(client.Url("/api/"))
	// Output: https://api.shapeways.com/api/v1/
}

func ExampleClient_GetApiInfo() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	// set users token/secret
	client.OauthCredentials.Token = "TOKEN"
	client.OauthCredentials.Secret = "SECRET"
	results, err := client.GetApiInfo()
}

func ExampleClient_GetCart() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	// set users token/secret
	client.OauthCredentials.Token = "TOKEN"
	client.OauthCredentials.Secret = "SECRET"
	results, err := client.GetCart()
}

func ExampleClient_GetCategories() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	// set users token/secret
	client.OauthCredentials.Token = "TOKEN"
	client.OauthCredentials.Secret = "SECRET"
	results, err := client.GetCategories()
}

func ExampleClient_GetCategory() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	// set users token/secret
	client.OauthCredentials.Token = "TOKEN"
	client.OauthCredentials.Secret = "SECRET"
	results, err := client.GetCategory(7)
}

func ExampleClient_GetMaterial() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	// set users token/secret
	client.OauthCredentials.Token = "TOKEN"
	client.OauthCredentials.Secret = "SECRET"
	results, err := client.GetMaterial(25)
}

func ExampleClient_GetMaterials() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	// set users token/secret
	client.OauthCredentials.Token = "TOKEN"
	client.OauthCredentials.Secret = "SECRET"
	results, err := client.GetMaterials()
}

func ExampleClient_GetModel() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	// set users token/secret
	client.OauthCredentials.Token = "TOKEN"
	client.OauthCredentials.Secret = "SECRET"
	results, err := client.GetModel(1670823)
}

func ExampleClient_GetModels() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	// set users token/secret
	client.OauthCredentials.Token = "TOKEN"
	client.OauthCredentials.Secret = "SECRET"
	results, err := client.GetModels(1)
}

func ExampleClient_GetModelInfo() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	// set users token/secret
	client.OauthCredentials.Token = "TOKEN"
	client.OauthCredentials.Secret = "SECRET"
	results, err := client.GetModelInfo(1670823)
}

func ExampleClient_GetModelFile() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	// set users token/secret
	client.OauthCredentials.Token = "TOKEN"
	client.OauthCredentials.Secret = "SECRET"

	// do not include raw file in result
	results, err := client.GetModelFile(1670823, 1, false)
	// include raw file in result
	results, err := client.GetModelFile(1670823, 1, true)
}

func ExampleClient_GetPrice() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	// set users token/secret
	client.OauthCredentials.Token = "TOKEN"
	client.OauthCredentials.Secret = "SECRET"

	priceData := url.Values{}
	priceData.Set("volume", "0.005")
	priceData.Set("area", "0.05")
	priceData.Set("xBoundMin", "0.1")
	priceData.Set("xBoundMax", "0.5")
	priceData.Set("yBoundMin", "0.1")
	priceData.Set("yBoundMax", "0.5")
	priceData.Set("zBoundMin", "0.1")
	priceData.Set("zBoundMax", "0.5")
	results, err := client.GetPrice(priceData)
}

func ExampleClient_GetPrinters() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	// set users token/secret
	client.OauthCredentials.Token = "TOKEN"
	client.OauthCredentials.Secret = "SECRET"
	results, err := client.GetPrinters()
}

func ExampleClient_GetPrinter() {
	client := shapeways.NewClient(
		"MY CONSUMER KEY",
		"MY CONSUMER SECRET",
		"http://localhost/callback",
	)
	// set users token/secret
	client.OauthCredentials.Token = "TOKEN"
	client.OauthCredentials.Secret = "SECRET"
	results, err := client.GetPrinter(1)
}
