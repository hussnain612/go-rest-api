package routes

import (
	"log"
	"net/http"

	"go-rest-api/controller"

	"github.com/gorilla/mux"
)

func Routes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.HomeLink)
	router.HandleFunc("/event", controller.CreateEvent).Methods("POST")
	router.HandleFunc("/event", controller.GetAllEvents).Methods("GET")
	router.HandleFunc("/event/{id}", controller.GetEvent).Methods("GET")
	router.HandleFunc("/event/{id}", controller.DeleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
