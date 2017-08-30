package main

import (
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func main() {

	router := mux.NewRouter()

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	intialRoutes(router, session)

	http.ListenAndServe(":3001", router)
}
