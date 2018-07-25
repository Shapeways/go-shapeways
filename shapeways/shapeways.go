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
  "encoding/base64"
  "strconv"
  "fmt"
  "errors"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

// Structs for holding return types
type Material struct {
  MaterialId string `json:materialId`
  Title string `json:`
}

type MaterialsMap struct {
  Result string `json.result`
  Materials map[string]Material `json.materials`
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


func (client *Oauth2Client) DoHttpRequest(req *http.Request) (*http.Response, error) {
  http_client := &http.Client{}
  resp, err := http_client.Do(req)
  check(err)
  if resp.StatusCode != 200 {
    fmt.Println("Request returned status " + resp.Status)
    return resp, errors.New("Request returned status " + resp.Status)
  }
  return resp, nil
}

func (client *Oauth2Client) UploadModel(Filename string) (string, error) {
  file_data, err := ioutil.ReadFile(Filename)
  check(err)
  enc_data := base64.StdEncoding.EncodeToString(file_data)
  // fmt.Println(enc_data)

  type ModelData struct {
    File string `json:"file"`
    FileName string `json:"fileName"`
    AcceptTermsAndConditions string `json:"acceptTermsAndConditions"`
    Description string `json:"description"`
    HasRightsToModel string `json:"hasRightsToModel"`
  }
  md := &ModelData {
      File: enc_data,
      FileName: "hi",
      AcceptTermsAndConditions: "1",
      HasRightsToModel: "1",
      Description: "Someone call a doctor, because this cube is SIIIICK.",
    }
  bytesToUpload, err := json.Marshal(md)
  check(err)
  fmt.Println(string(bytesToUpload))
  // req, err := http.NewRequest("POST", "https://api.shapeways.com/models/v1", bytes.NewBuffer(bytesToUpload))
  // req, err := http.NewRequest("POST", "https://api.jw.nyc.shapeways.net/models/v1", bytes.NewBuffer(bytesToUpload))
  req, err := http.NewRequest("POST", "http://requestbin.fullcontact.com/y5887gy5", bytes.NewBuffer(bytesToUpload))
  // // req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Authorization","Bearer " + client.BearerToken)

  resp, err := client.DoHttpRequest(req)
  defer resp.Body.Close()
  type UploadResult struct {
    Result string `json.result`
  }
  var result UploadResult
  json.NewDecoder(resp.Body).Decode(&result)
  fmt.Println(result)
  return "", nil
}

// Authorize using Oauth2
func (client *Oauth2Client) Authenticate() (string, error) {
  // Result struct for json decode
  type Result struct {
      Access_token string `json.access_token`
      Expiration_time int `json.expires_in`
      Token_type string `json.token_type`
  }

  // Build request, including headers, auth type, and post data
  var jsonStr = []byte("grant_type=client_credentials")
  req, err := http.NewRequest("POST", "https://api.shapeways.com/oauth2/token", bytes.NewBuffer(jsonStr))
  // req, err := http.NewRequest("POST", "http://api.jw.nyc.shapeways.net/oauth2/token", bytes.NewBuffer(jsonStr))
  check(err)
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  req.SetBasicAuth(client.ConsumerKey, client.ConsumerSecret)
  resp, err := client.DoHttpRequest(req)
  if err != nil {
    return "failure", errors.New("Request failed")
  }
  defer resp.Body.Close()
  var result Result
  json.NewDecoder(resp.Body).Decode(&result)
  if result.Access_token == "" {
    panic("Access token not returned!")
  }
  client.BearerToken = result.Access_token
  return "success", nil
}

func (client *Oauth2Client) GetMaterials() (string, error) {
  /* Get a specific material in our portfolio */

  // Build request
  var materials_url = "https://api.shapeways.com/materials/v1"
  req, err := http.NewRequest("GET", materials_url, nil)
  req.Header.Set("Authorization","Bearer " + client.BearerToken)
  resp, err := client.DoHttpRequest(req)
  if err != nil {
		panic(err)
	}
  defer resp.Body.Close()

  var data MaterialsMap
	json.NewDecoder(resp.Body).Decode(&data)
  return "success", nil
}

func (client *Oauth2Client) GetMaterial(MaterialId int) (Material, error) {
  /* Get all materials in our portfolio */

  // Build request
  var materials_url = "https://api.shapeways.com/materials/" + strconv.Itoa(MaterialId) + "/v1"
  req, err := http.NewRequest("GET", materials_url, nil)
  req.Header.Set("Authorization","Bearer " + client.BearerToken)

  resp, err := client.DoHttpRequest(req)
  if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
  var data Material
  json.NewDecoder(resp.Body).Decode(&data)
  return data, nil
}
