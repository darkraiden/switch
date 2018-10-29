package handlers

import (
	"fmt"
	"net/http"
)

func NewHTTPHandler(paths map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := paths[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fmt.Fprintf(w, "404: Not found")
	}
}
