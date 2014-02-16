package server

import "github.com/gorilla/mux"

func RegisterHandlers() {
	r := mux.NewRouter()
}

// badRequest is handled by setting the status code in
// reply to StatusBadRequest
type badRequest struct{ error }

// notFound is handled by setting the status code in
// reply to StatusNotFound
type notFound struct{ error }

// TODO fill in routes
