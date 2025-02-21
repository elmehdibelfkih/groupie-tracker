package internal

import (
	"groupie-tracker/error"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		error.MethodNotAllowed(w, r)
		return
	}
	if r.URL.Path != "/" {
		error.NotFoundError(w, r)
		return
	}
	INDEX_TMPL.Execute(w, ArtistsStruct)
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		error.MethodNotAllowed(w, r)
		return
	}
	filePath, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		if os.IsNotExist(err) {
			error.NotFoundError(w, r)
			return
		}
		error.InternalServerError(w, r)
		return
	}
	if filePath.IsDir() {
		error.NotFoundError(w, r)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		error.MethodNotAllowed(w, r)
		return
	}
	tmp := strings.Split(r.URL.Path, "/")
	if len(tmp) != 3 {
		error.BadRequest(w, r)
		return
	}
	id, err := strconv.ParseUint(tmp[2], 10, 0)
	if err != nil || id < 1 || id > 52 {
		error.BadRequest(w, r)
		return
	}
	var location Location
	FetchAPI(ArtistsStruct[id-1].Locations, &location)
	location.Artist = &ArtistsStruct[id-1]
	LOCATIONS_TMPL.Execute(w, location)
}

func DateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		error.MethodNotAllowed(w, r)
		return
	}
	tmp := strings.Split(r.URL.Path, "/")
	if len(tmp) != 3 {
		error.BadRequest(w, r)
		return
	}
	id, err := strconv.ParseUint(tmp[2], 10, 0)
	if err != nil || id < 1 || id > 52 {
		error.BadRequest(w, r)
		return
	}
	var date Date
	FetchAPI(ArtistsStruct[id-1].ConcertDates, &date)
	date.Artist = &ArtistsStruct[id-1]
	DATES_TMPL.Execute(w, date)
}

func RelationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		error.MethodNotAllowed(w, r)
		return
	}
	tmp := strings.Split(r.URL.Path, "/")
	if len(tmp) != 3 {
		error.BadRequest(w, r)
		return
	}
	id, err := strconv.ParseUint(tmp[2], 10, 0)
	if err != nil || id < 1 || id > 52 {
		error.BadRequest(w, r)
		return
	}
	var relation Relation
	FetchAPI(ArtistsStruct[id-1].Relations, &relation)
	relation.Artist = &ArtistsStruct[id-1]
	RELATIONS_TMPL.Execute(w, relation)
}
