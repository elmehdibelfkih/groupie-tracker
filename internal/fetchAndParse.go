package internal

import (
	"encoding/json"
	"groupie-tracker/error"
	"io"
	"net/http"
)

var (
	APIsStruct    APIs
	ArtistsStruct Artists
)

func FetchAPIs(url string, w http.ResponseWriter) {

	response, err := http.Get(url)
	// TODO: check the response status
	if err != nil {
		// TODO: INTERNEL SERVER ERORR !
		error.BadGateway(w)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		// TODO: INTERNEL SERVER ERORR !
		error.BadGateway(w)
		return
	}
	err1 := json.Unmarshal(body, &APIsStruct)
	if err1 != nil {
		error.BadGateway(w)
		return
	}
	FetchAPI(APIsStruct.Artists, &ArtistsStruct, w)
}

func FetchAPI(url string, v any, w http.ResponseWriter) {

	response, err := http.Get(url)
	if err != nil {
		error.BadGateway(w)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		error.BadGateway(w)
		return
	}
	err1 := json.Unmarshal(body, v)
	if err1 != nil {
		error.BadGateway(w)
		return
	}
}
