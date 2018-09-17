package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"io"
	"fmt"
	"encoding/json"
	"io/ioutil"
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
	io.WriteString(w, "Event list")
}

// create api
func (api *ApiController) createEvent(w http.ResponseWriter, r *http.Request) {

	eventDto := EventDto{}

	if err := requestToDto(r, &eventDto); err != nil {
		println(err.Error())
	}

	fmt.Printf("%+v\n", eventDto)
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

func requestToDto(r *http.Request, dto interface{}) error {
	var body []byte
	var err error

	if body, err = ioutil.ReadAll(r.Body); err != nil {
		return err
	}

	return json.Unmarshal(body, dto)
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
