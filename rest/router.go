package rest

import "github.com/gorilla/mux"

func createRouter() *mux.Router {
	r := mux.NewRouter()

	r.Handle("/", makeHandler((*handler).handleRoot))
	r.Handle("/auth", makeHandler((*handler).handleAuth)).Methods("POST")

	return r
}
