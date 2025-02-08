package main

import (
	"fmt"
	"groupie-tracker/internal"
	"net/http"
	"os"
)

func main() {
	internal.FetchAPIs(internal.API_URL)
	http.HandleFunc("/", internal.RootHandler)
	fmt.Println("Server is running on http://localhost:8080")
	err1 := http.ListenAndServe(":8080", nil)
	if err1 != nil {
		fmt.Fprintln(os.Stderr, err1.Error())
	}
}
