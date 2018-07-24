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
	"net/http"
  "io/ioutil"
  "bytes"
  "encoding/json"
  "strconv"
  "fmt"
)
// Structs for holding return types
type Material struct {
  MaterialId string
  Title string

}

// Setup a new shapeways.Client, use this instead of creating a Client directly
func NewClient(ConsumerKey string, ConsumerSecret string) Oauth2Client {
	client := Oauth2Client{
		BaseUrl:          "https://api.shapeways.com",
		APIVersion:       "v1",
    ConsumerKey:      ConsumerKey,
    ConsumerSecret:   ConsumerSecret,
	}
  return client
}

// Represents a shapeways API Client
type Oauth2Client struct {
	BaseUrl, APIVersion          string
	ConsumerKey, ConsumerSecret  string
  BearerToken                  string
}

// Authorize using Oauth2
func (client *Oauth2Client) Authenticate() (string, error) {
  // Result struct for json decode
  type Result struct {
      Access_token string
  }

  // Build request, including headers, auth type, and post data
  var jsonStr = []byte("grant_type=client_credentials")
  req, err := http.NewRequest("POST", "https://api.shapeways.com/oauth2/token", bytes.NewBuffer(jsonStr))
  if err != nil {
		panic(err)
	}
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  req.SetBasicAuth(client.ConsumerKey, client.ConsumerSecret)

  // Create client and make request
  http_client := &http.Client{}
  resp, err := http_client.Do(req)
  if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

  // Process request body
  body, err := ioutil.ReadAll(resp.Body)
  var result Result
  err = json.Unmarshal(body, &result)
  if err != nil {
    panic(err)
  }
  client.BearerToken = result.Access_token
  // fmt.Println(client.BearerToken)
  return "success", nil
}

func (client *Oauth2Client) GetMaterial(MaterialId int) (Material, error) {


  // Build request
  var materials_url = "https://api.shapeways.com/materials/" + strconv.Itoa(MaterialId) + "/v1"

  fmt.Println(materials_url)
  req, err := http.NewRequest("GET", materials_url, nil)
  req.Header.Set("Authorization","Bearer " + client.BearerToken)

  // Execute request
  http_client := &http.Client{}
  resp, err := http_client.Do(req)
  if err != nil {
    panic(err)
	}
	defer resp.Body.Close()

  // Process request body
  body, err := ioutil.ReadAll(resp.Body)
  var material Material
  err = json.Unmarshal(body, &material)
  if err != nil {
    panic(err)
  }
  return material, nil
}
