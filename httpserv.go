package main

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/couchbaselabs/logg"
)

func handleEndpoint(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, readHTMLFile("greet.html"))
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		response := authenticate(r)
		logg.LogTo(logTag, "%s", response)
		io.WriteString(w, response)
		break
	default:
		io.WriteString(w, readHTMLFile("greet.html"))
	}
}

func readHTMLFile(file string) string {
	b, err := ioutil.ReadFile("html/" + file)

	if err != nil {
		logg.LogPanic("Error reading static file %s : %v", file, err)
	}

	return string(b)
}
