package main

import "net/http"

func handleError(w http.ResponseWriter, errors []error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
}

func EventList(w http.ResponseWriter, r *http.Request) {

}

func GetEvent(w http.ResponseWriter, r *http.Request) {

}

func CreateEvent(w http.ResponseWriter, r *http.Request) {

}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {

}
