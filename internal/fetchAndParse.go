package internal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
)

var APIsStruct APIs
var ArtistsStruct Artists

// var LocationsStruct Locations
// var DatesStruct Dates
var RelationsStruct Relations

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
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go FetchAPI(APIsStruct.Artists, &ArtistsStruct, wg)
	// go FetchAPI(APIsStruct.Locations, &LocationsStruct, wg)
	// go FetchAPI(APIsStruct.Dates, &DatesStruct, wg)
	go FetchAPI(APIsStruct.Relation, &RelationsStruct, wg)
	wg.Wait()
	MappingData()
	println("================")
	println(ArtistsStruct[0].Relation.DatesLocations["dunedin-new_zealand"][0])
	println("================")
}

func FetchAPI(url string, v any, wg *sync.WaitGroup) {
	defer wg.Done()
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

func MappingData() {
	for i := range ArtistsStruct {
		ArtistsStruct[i].Relation = &RelationsStruct.Relations[i]
	}
}
