package internal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var APIsStruct APIs
var ArtistsStruct Artists

func FetchAPIs(url string) {

	response, err := http.Get(url)
	// TODO: check the response status
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err1 := json.Unmarshal(body, &APIsStruct)
	if err1 != nil {
		log.Fatal(err)
	}
	FetchAPI(APIsStruct.Artists, &ArtistsStruct)
}

func FetchAPI(url string, v any) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		log.Fatal(err)
	}
}
