package internal

// import "json"

const API_URL = "https://groupietrackers.herokuapp.com/api"

type APIs struct {
	Artists   string `jason:"artists"`
	Locations string `jason:"locations"`
	Dates     string `jason:"dates"`
	Relation  string `jason:"relation"`
}

// Artists
type Artists []Artist

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	// Locations    string   `json:"locations"`    // FIXME: must be a pointer to a struct
	// ConcertDates string   `json:"concertDates"` // FIXME: must be a pointer to a struct
	// Relations    string   `json:"relations"`    // FIXME: must be a pointer to a struct
	Relation *Relation
}

// // Locations
// type Locations struct {
// 	Location []Location `json:"index"`
// }

// type Location struct {
// 	Id        int      `json:"id"`
// 	Locations []string `json:"locations"`
// 	Dates     string   `json:"dates"` // FIXME: must be a struct
// }

// // Dates
// type Dates struct {
// 	Dates []Date `json:"index"`
// }

// type Date struct {
// 	Id    int      `json:"id"`
// 	Dates []string `json:"dates"`
// }

// Relations
type Relations struct {
	Relations []Relation `json:"index"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
