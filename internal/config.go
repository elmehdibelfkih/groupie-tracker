package internal

import "html/template"

const API_URL = "https://groupietrackers.herokuapp.com/api"

var INDEX_TMPL *template.Template
var LOCATIONS_TMPL *template.Template
var DATES_TMPL *template.Template
var RELATIONS_TMPL *template.Template

const (
	INDEX_PATH     = "./templates/index.html"
	LOCATIONS_PATH = "./templates/locations.html"
	DATES_PATH     = "./templates/dates.html"
	RELATIONS_PATH = "./templates/relations.html"
)

type APIs struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Artists []Artist
type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type Location struct {
	Id        int      `json:"id"`
	Artist    *Artist  `json:"-"`
	Locations []string `json:"locations"`
}
type Date struct {
	Id     int      `json:"id"`
	Artist *Artist  `json:"-"`
	Dates  []string `json:"dates"`
}

type Relation struct {
	Id             int                 `json:"id"`
	Artist         *Artist             `json:"-"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
