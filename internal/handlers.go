package internal

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"groupie-tracker/error"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodGet {
		error.MethodNotAllowed(w)
		return
	}
	if r.URL.Path != "/" {
		error.NotFoundError(w)
		return
	}
	FetchAPIs(API_URL, w)
	INDEX_TMPL.Execute(w, ArtistsStruct)
}


func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		error.MethodNotAllowed(w)
		return
	}
	filePath, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		if os.IsNotExist(err) {
			error.NotFoundError(w)
			return
		}
		error.InternalServerError(w)
		return
	}
	if filePath.IsDir() {
		error.NotFoundError(w)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		error.MethodNotAllowed(w)
		return
	}
	tmp := strings.Split(r.URL.Path, "/")
	if len(tmp) != 3 {
		error.BadRequest(w)
		return
	}
	id, err := strconv.ParseUint(tmp[2], 10, 0)
	if err != nil || id < 1 || id > 52 {
		error.BadRequest(w)
		return
	}
	var location Location
	FetchAPI(ArtistsStruct[id-1].Locations, &location, w)
	location.Artist = &ArtistsStruct[id-1]
	LOCATIONS_TMPL.Execute(w, location)
}

func DateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		error.MethodNotAllowed(w)
		return
	}
	tmp := strings.Split(r.URL.Path, "/")
	if len(tmp) != 3 {
		error.BadRequest(w)
		return
	}
	id, err := strconv.ParseUint(tmp[2], 10, 0)
	if err != nil || id < 1 || id > 52 {
		error.BadRequest(w)
		return
	}
	var date Date
	FetchAPI(ArtistsStruct[id-1].ConcertDates, &date, w)
	date.Artist = &ArtistsStruct[id-1]
	// TODO: trim *
	date.Dates = Removeandreturn(date.Dates)
	DATES_TMPL.Execute(w, date)
}

func RelationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		error.MethodNotAllowed(w)
		return
	}
	tmp := strings.Split(r.URL.Path, "/")
	if len(tmp) != 3 {
		error.BadRequest(w)
		return
	}
	id, err := strconv.ParseUint(tmp[2], 10, 0)
	if err != nil || id < 1 || id > 52 {
		error.BadRequest(w)
		return
	}
	var relation Relation
	FetchAPI(ArtistsStruct[id-1].Relations, &relation, w)
	relation.Artist = &ArtistsStruct[id-1]
	RELATIONS_TMPL.Execute(w, relation)
}
