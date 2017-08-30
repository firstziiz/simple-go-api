package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Room struct {
	ID          bson.ObjectId `bson:"_id"`
	Code        string        `bson:"code"`
	Title       string        `bson:"title"`
	OpenSending bool          `bson:"openSending"`
	UpdatedAt   time.Time     `bson:"updatedAt"`
	CreatedAt   time.Time     `bson:"createdAt"`
	Imgs        Img           `bson:"imgs"`
}

type Img struct {
	Cover string `bson:"cover"`
}

func GetAllRooms(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		session := s.Copy()
		defer session.Close()

		c := session.DB("alchemist_ask").C("rooms")

		var rooms []Room
		err := c.Find(bson.M{}).All(&rooms)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all rooms: ", err)
			return
		}

		respBody, err := json.MarshalIndent(rooms, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}
