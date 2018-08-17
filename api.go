package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"io"
	"fmt"
	"time"
)

type ApiController struct {
	router *mux.Router
}

// get list api response
func (api *ApiController) eventList(w http.ResponseWriter, r *http.Request) {
	println("Route started")

	io.WriteString(w, "Events list")

	time.Sleep(time.Duration(5) * time.Second)

	println("Route stopped")
}

// create api
func (api *ApiController) createEvent(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Create event")
}

// get the event
func (api *ApiController) getEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	io.WriteString(w, fmt.Sprintf("Get event %s", vars["id"]))
}

// delete the event
func (api *ApiController) deleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	io.WriteString(w, fmt.Sprintf("Delete event %s", vars["id"]))
}

func NewApiController() *ApiController {
	controller := &ApiController{
		router: mux.NewRouter(),
	}

	controller.router.HandleFunc("/api/v1/event", controller.createEvent).Methods("POST")
	controller.router.HandleFunc("/api/v1/event", controller.eventList).Methods("GET")
	controller.router.HandleFunc("/api/v1/event/{id}", controller.getEvent).Methods("GET")
	controller.router.HandleFunc("/api/v1/event/{id}", controller.deleteEvent).Methods("DELETE")

	return controller
}
