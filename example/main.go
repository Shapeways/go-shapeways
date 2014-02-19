package main

import (
	"encoding/json"
	"github.com/Shapeways/go-shapeways/shapeways"
	"log"
	"net/http"
)

var client = shapeways.NewClient(
	"CONSUMER KEY",
	"CONSUMER SECRET",
	"http://localhost:3000/callback",
)

func handleHome(res http.ResponseWriter, req *http.Request) {
	data, err := client.GetApiInfo()
	if err != nil {
		log.Fatal("Error Getting Request Token: ", err)
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data)
}

func handleLogin(res http.ResponseWriter, req *http.Request) {
	url, err := client.Connect()
	if err != nil {
		log.Fatal("Error Getting Request Token: ", err)
	} else {
		http.Redirect(res, req, url, 302)
	}
}

func handleCallback(res http.ResponseWriter, req *http.Request) {
	err := client.VerifyURL(req.URL)
	if err != nil {
		log.Fatal("Error Getting Request Token: ", err)
	}
	http.Redirect(res, req, "/", 302)
}

func main() {
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	http.HandleFunc("/", handleHome)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Could not bind: ", err)
	}
}
