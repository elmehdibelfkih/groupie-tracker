package internal

import (
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// TODO: return MethodNotAllowed template
		return
	}

}
