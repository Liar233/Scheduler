package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"io"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type EventDto struct {
	ID       string `json:"id"`
	Channel  string `json:"channel"`
	Payload  string `json:"payload"`
	FireTime string `json:"fireTime"`
}

type ApiController struct {
	router *mux.Router
}

// get list api response
func (api *ApiController) eventList(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "List event")
}

// create api
func (api *ApiController) createEvent(w http.ResponseWriter, r *http.Request) {
	eventDto :=	EventDto{}

	if err := jsonToDto(r, &eventDto); err != nil {
		println(err.Error())
	}

	io.WriteString(w, fmt.Sprintf("%+v\n", eventDto))
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

// try to transform request body from json to struct
func jsonToDto(r *http.Request, dto interface{}) error {
	content, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(content, dto)
}
