// The MIT License (MIT) Copyright (c) 2014 Shapeways <api@shapeways.com> (http://developers.shapeways.com)

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
// PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF
// CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// The shapeways package is a client library for the Shapeways API http://developers.shapeways.com
package shapeways

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/brettlangdon/go-oauth/oauth"
	"net/http"
	"net/url"
	"strings"
)

// Setup a new shapeways.Client, use this instead of creating a Client directly
func NewClient(ConsumerKey string, ConsumerSecret string, CallbackUrl string) Client {
	client := Client{
		BaseUrl:          "https://api.shapeways.com",
		APIVersion:       "v1",
		CallbackUrl:      CallbackUrl,
		OauthCredentials: &oauth.Credentials{},
	}

	client.OauthClient = oauth.Client{
		TemporaryCredentialRequestURI: client.Url("/oauth1/request_token/"),
		TokenRequestURI:               client.Url("/oauth1/access_token/"),
		Credentials: oauth.Credentials{
			Token:  ConsumerKey,
			Secret: ConsumerSecret,
		},
	}

	return client
}

// Represents a shapeways API Client
type Client struct {
	BaseUrl, APIVersion     string
	OauthToken, OauthSecret string
	CallbackUrl             string
	OauthClient             oauth.Client
	OauthCredentials        *oauth.Credentials
}

func (client *Client) get(url string, form url.Values) (interface{}, error) {
	resp, err := client.OauthClient.Get(http.DefaultClient, client.OauthCredentials, url, form)
	defer resp.Body.Close()
	var data interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	return data, err
}

func (client *Client) delete(url string, form url.Values) (interface{}, error) {
	resp, err := client.OauthClient.Delete(http.DefaultClient, client.OauthCredentials, url, form)
	defer resp.Body.Close()
	var data interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	return data, err
}

func (client *Client) post(url string, form url.Values) (interface{}, error) {
	resp, err := client.OauthClient.Post(http.DefaultClient, client.OauthCredentials, url, form)
	defer resp.Body.Close()
	var data interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	return data, err
}

func (client *Client) put(url string, form url.Values) (interface{}, error) {
	resp, err := client.OauthClient.Put(http.DefaultClient, client.OauthCredentials, url, form)
	defer resp.Body.Close()
	var data interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	return data, err
}

// Request temporary credentials and get an authorization url
// to redirect the user to for authentication
func (client *Client) Connect() (string, error) {
	tempCreds, authUrl, err := client.OauthClient.RequestTemporaryCredentials(http.DefaultClient, client.CallbackUrl, nil)
	if err != nil {
		return "", err
	}
	client.OauthCredentials = tempCreds
	return authUrl, nil
}

// Verify the token and verifier from the authroization callback
func (client *Client) Verify(token string, verifier string) error {
	client.OauthCredentials.Token = token
	creds, _, err := client.OauthClient.RequestToken(http.DefaultClient, client.OauthCredentials, verifier)
	if err != nil {
		return err
	}
	client.OauthCredentials = creds
	return nil
}

// parse the callback url for the token and verifier and call Client.Verify
func (client *Client) VerifyURL(url *url.URL) error {
	params := url.Query()
	token := ""
	verifier := ""
	if len(params["oauth_token"]) > 0 && params["oauth_token"][0] != "" {
		token = params["oauth_token"][0]
	} else {
		return errors.New("Missing or empty 'oauth_token' parameter")
	}
	if len(params["oauth_verifier"]) > 0 && params["oauth_verifier"][0] != "" {
		verifier = params["oauth_verifier"][0]
	} else {
		return errors.New("Missing or empty 'oauth_verifier' parameter")
	}

	return client.Verify(token, verifier)
}

// generate the proper full url from the BaseUrl and APIVersion given the path portion
func (client *Client) Url(path string) string {
	path = strings.Trim(path, "/")
	baseUrl := strings.Trim(client.BaseUrl, "/")
	version := strings.Trim(client.APIVersion, "/")
	return fmt.Sprintf("%s/%s/%s", baseUrl, path, version)
}

func (client *Client) GetApiInfo() (interface{}, error) {
	return client.get(client.Url("/api"), url.Values{})
}

func (client *Client) GetCart() (interface{}, error) {
	return client.get(client.Url("/orders/cart/"), url.Values{})
}

func (client *Client) GetMaterial(MaterialId int) (interface{}, error) {
	return client.get(client.Url(fmt.Sprintf("/materials/%d/", MaterialId)), url.Values{})
}

func (client *Client) GetMaterials() (interface{}, error) {
	return client.get(client.Url("/materials/"), url.Values{})
}

func (client *Client) GetModel(ModelId int) (interface{}, error) {
	return client.get(client.Url(fmt.Sprintf("/models/%d/", ModelId)), url.Values{})
}

func (client *Client) AddModel() (interface{}, error) {
	required := []string{"file", "fileName"}
	// for key := range required {

	// }

	return required, nil
}
