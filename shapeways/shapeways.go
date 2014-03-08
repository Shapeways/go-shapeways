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
	"encoding/base64"
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

// Make an API call to GET /api/v1 https://developers.shapeways.com/docs?li=dh_docs#GET_-api-v1-1
func (client *Client) GetApiInfo() (interface{}, error) {
	return client.get(client.Url("/api"), url.Values{})
}

// Make an API call to GET /api/v1 https://developers.shapeways.com/docs?li=dh_docs#GET_-api-v1-1
func (client *Client) GetCart() (interface{}, error) {
	return client.get(client.Url("/orders/cart/"), url.Values{})
}

// Make an API call to POST /orders/cart/v1] https://developers.shapeways.com/docs?li=dh_docs#POST_-orders-cart-v1
func (client *Client) AddToCart(CartData url.Values) (interface{}, error) {
	if CartData.Get("modelId") == "" {
		return nil, errors.New("shapeways.Client.AddToCart missing required key: modelId")
	}

	return client.post(client.Url("/orders/cart/"), CartData)
}

// Make an API call to GET /materials/{materialId}/v1 https://developers.shapeways.com/docs?li=dh_docs#GET_-materials-materialId-v1
func (client *Client) GetMaterial(MaterialId int) (interface{}, error) {
	return client.get(client.Url(fmt.Sprintf("/materials/%d/", MaterialId)), url.Values{})
}

// Make an API call to GET /materials/v1 https://developers.shapeways.com/docs?li=dh_docs#GET_-materials-v1
func (client *Client) GetMaterials() (interface{}, error) {
	return client.get(client.Url("/materials/"), url.Values{})
}

// Make an API call to GET /models/v1 https://developers.shapeways.com/docs?li=dh_docs#GET_-models-v1
func (client *Client) GetModels(Page int) (interface{}, error) {
	params := url.Values{}
	params.Set("page", fmt.Sprint("%d", Page))
	return client.get(client.Url("/models/"), params)
}

// Make an API call to GET /models/{modelId}/v1 https://developers.shapeways.com/docs?li=dh_docs#GET_-models-modelId-v1
func (client *Client) GetModel(ModelId int) (interface{}, error) {
	return client.get(client.Url(fmt.Sprintf("/models/%d/", ModelId)), url.Values{})
}

// Make an API call to POST /models/v1 https://developers.shapeways.com/docs?li=dh_docs#POST_-models-v1
func (client *Client) AddModel(ModelData url.Values) (interface{}, error) {
	required := []string{"file", "fileName", "acceptTermsAndConditions", "hasRightsToModel"}
	for _, key := range required {
		if ModelData.Get(key) == "" {
			return nil, errors.New(
				fmt.Sprintf("shapeways.Client.AddModel missing required key: %s", key),
			)
		}
	}

	ModelData.Set(
		"file",
		url.QueryEscape(
			base64.StdEncoding.EncodeToString([]byte(ModelData.Get("file"))),
		),
	)
	return client.post(client.Url("/models/"), ModelData)
}

// Make an API call to GET /models/{modelId}/info/v1 https://developers.shapeways.com/docs?li=dh_docs#GET_-models-modelId-info-v1
func (client *Client) GetModelInfo(ModelId int) (interface{}, error) {
	return client.get(client.Url(fmt.Sprintf("/models/%d/info/", ModelId)), url.Values{})
}

// Make an API call to PUT /models/{modelId}/info/v1 https://developers.shapeways.com/docs?li=dh_docs#PUT_-models-modelId-info-v1
func (client *Client) UpdateModelInfo(ModelId int, ModelData url.Values) (interface{}, error) {
	return client.put(client.Url(fmt.Sprintf("/models/%d/", ModelId)), ModelData)
}

// Make an API call to DELETE /models/{modelId}/v1 https://developers.shapeways.com/docs?li=dh_docs#DELETE_-models-modelId-v1
func (client *Client) DeleteModel(ModelId int) (interface{}, error) {
	return client.delete(client.Url(fmt.Sprintf("/models/%d/", ModelId)), url.Values{})
}

// Make an API call to GET /models/{modelId}/files/{fileVersion}/v1 https://developers.shapeways.com/docs?li=dh_docs#GET_-models-modelId-files-fileVersion-v1
func (client *Client) GetModelFile(ModelId int, FileVersion int, IncludeFile bool) (interface{}, error) {
	params := url.Values{}
	if IncludeFile {
		params.Set("file", "1")
	} else {
		params.Set("file", "0")
	}
	return client.get(client.Url(fmt.Sprintf("/models/%d/files/%d/", ModelId, FileVersion)), params)
}

// Make an API call to POST /models/{modelId}/photos/v1 https://developers.shapeways.com/docs?li=dh_docs#POST_-models-modelId-photos-v1
func (client *Client) AddModelPhoto(ModelId int, PhotoData url.Values) (interface{}, error) {
	if PhotoData.Get("file") == "" {
		return nil, errors.New("shapeways.Client.AddModelPhoto missing required key: file")
	}

	PhotoData.Set(
		"file",
		url.QueryEscape(
			base64.StdEncoding.EncodeToString([]byte(PhotoData.Get("file"))),
		),
	)
	return client.post(client.Url(fmt.Sprintf("/models/%d/", ModelId)), PhotoData)
}

// Make an API call to POST /models/{modelId}/files/v1 https://developers.shapeways.com/docs?li=dh_docs#POST_-models-modelId-files-v1
func (client *Client) AddModelFile(ModelId int, FileData url.Values) (interface{}, error) {
	required := []string{"file", "fileName", "acceptTermsAndConditions", "hasRightsToModel"}
	for _, key := range required {
		if FileData.Get(key) == "" {
			return nil, errors.New(
				fmt.Sprintf("shapeways.Client.AddModelFile missing required key: %s", key),
			)
		}
	}

	FileData.Set(
		"file",
		url.QueryEscape(
			base64.StdEncoding.EncodeToString([]byte(FileData.Get("file"))),
		),
	)
	return client.post(client.Url(fmt.Sprintf("/models/%d/files/", ModelId)), FileData)
}

// Make an API call to GET /printers/v1 https://developers.shapeways.com/docs?li=dh_docs#GET_-printers-v1
func (client *Client) GetPrinters() (interface{}, error) {
	return client.get(client.Url("/printers/"), url.Values{})
}

// Make an API call to GET /printers/{printerId}/v1 https://developers.shapeways.com/docs?li=dh_docs#GET_-printers-printerId-v1
func (client *Client) GetPrinter(PrinterId int) (interface{}, error) {
	return client.get(client.Url(fmt.Sprintf("/printers/%d/", PrinterId)), url.Values{})
}

// Make an API call GET /categories/v1 https://developers.shapeways.com/docs?li=dh_docs#GET_-categories-v1
func (client *Client) GetCategories() (interface{}, error) {
	return client.get(client.Url("/categories/"), url.Values{})
}

// Make an API call to GET /categories/{categoryId}/v1 https://developers.shapeways.com/docs?li=dh_docs#GET_-categories-categoryId-v1
func (client *Client) GetCategory(CatId int) (interface{}, error) {
	return client.get(client.Url(fmt.Sprintf("/categories/%d/", CatId)), url.Values{})
}

// Make an API call to POST /price/v1 https://developers.shapeways.com/docs?li=dh_docs#POST_-price-v1
func (client *Client) GetPrice(PriceData url.Values) (interface{}, error) {
	required := []string{
		"volume", "area", "xBoundMin", "xBoundMax",
		"yBoundMin", "yBoundMax", "zBoundMin", "zBoundMax",
	}
	for _, key := range required {
		if PriceData.Get(key) == "" {
			return nil, errors.New(
				fmt.Sprintf("shapeways.Client.GetPrice missing required key: %s", key),
			)
		}
	}
	return client.post(client.Url("/price//"), PriceData)
}
