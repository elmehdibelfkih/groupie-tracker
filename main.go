package main

import (
	"fmt"
	"groupie-tracker/internal"
	"net/http"
	"os"
)

func main() {
	internal.FetchAPIs(internal.API_URL)
	internal.InitTmpl()
	http.HandleFunc("/static/", internal.StaticHandler)
	http.HandleFunc("/", internal.RootHandler)
	http.HandleFunc("/Locations/{id}", internal.LocationHandler)
	http.HandleFunc("/Dates/{id}", internal.DateHandler)
	http.HandleFunc("/Relation/{id}", internal.RelationHandler)
	fmt.Println("Server is running on http://localhost:8080")
	err1 := http.ListenAndServe(":8080", nil)
	if err1 != nil {
		fmt.Fprintln(os.Stderr, err1.Error())
	}
}
