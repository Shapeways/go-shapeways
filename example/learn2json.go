package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Artist struct {
	Id             int
	Name           string
	Resource_url   string
	Releases_url   string
	Uri            string
	Realname       string
	Profile        string
	Data_quality   string
	Namevariations []string
	Aliases        []struct {
		Id           int
		Name         string
		Resource_url string
	}
	Urls   []string
	Images []struct {
		Type         string
		Width        int
		Height       int
		Uri          string
		Uri150       string
		Resource_url string
	}
}

func main() {
	//Query an artist.
	res, _ := http.Get("http://api.discogs.com/artists/1373")

	temp, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(temp))
	var artist Artist
	err := json.Unmarshal(temp, &artist)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	fmt.Println(artist.Name)
	// fmt.Println(artist.Profile)
	// fmt.Println(artist.Releases_url)
	// fmt.Println(artist.Namevariations)
	fmt.Println(artist.Images[1].Height)
}
