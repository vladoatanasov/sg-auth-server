package main

import (
	"io"
	"net/http"

	"github.com/couchbaselabs/logg"
)

func handleEndpoint(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{\"error\" : \"\"}")
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		response := authenticate(r)
		logg.LogTo(logTag, "%s", response)
		io.WriteString(w, response)
		break
	default:
		io.WriteString(w, "{\"error\" : \"\"}")
	}
}
