package main

import (
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func intialRoutes(router *mux.Router, session *mgo.Session) {
	router.HandleFunc("/rooms", GetAllRooms(session)).Methods("GET")
	// router.HandleFunc("/people/{id}", GetPersonEndpoint(session)).Methods("GET")
	// router.HandleFunc("/people/{id}", CreatePersonEndpoint(session)).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePersonEndpoint(session)).Methods("DELETE")
}
