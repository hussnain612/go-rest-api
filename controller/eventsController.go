package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go-rest-api/model"

	"github.com/gorilla/mux"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!")
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent model.Event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter data with Title and Description")
	}

	json.Unmarshal(reqBody, &newEvent)
	model.Events = append(model.Events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	evntId := mux.Vars(r)["id"]

	for _, event := range model.Events {
		if event.ID == evntId {
			json.NewEncoder(w).Encode(event)
		}
	}
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Events)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]

	for i, event := range model.Events {
		if event.ID == eventId {
			model.Events = append(model.Events[:i], model.Events[i+1:]...)
			fmt.Fprintf(w, "The event hase be deleted. ID: %v", eventId)
		}
	}
}
