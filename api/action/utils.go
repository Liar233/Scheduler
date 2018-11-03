package action

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type ApiResponse struct {
	IsSuccess bool `json:"isSuccess"`
}

type SuccessResponse struct {
	ApiResponse
	Payload interface{} `json:"payload"`
}

type ErrorResponse struct {
	ApiResponse
	Error string `json:"error"`
}

func FillHeaders(w http.ResponseWriter) {
	w.Header().Add("content-type", "application/json")
	w.Header().Add("cache-control", "no-cache, private")
}

func WriteErrorResponse(w http.ResponseWriter, err error) {
	response := ErrorResponse{
		Error: fmt.Sprintf("%v", err),
	}
	response.IsSuccess = false

	content, _ := json.Marshal(response)

	FillHeaders(w)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(content)
}

func WriteSuccessResponse(w http.ResponseWriter, payload interface{}) {
	response := SuccessResponse{
		Payload: payload,
	}
	response.IsSuccess = true

	content, _ := json.Marshal(response)

	FillHeaders(w)
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}
