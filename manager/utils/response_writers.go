package utils

import (
	"encoding/json"
	"net/http"
)

func PayloadResponse(w http.ResponseWriter, r *http.Request, payload interface{}) {
	// generic function to send an HTTP Response with payload
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

func NoRecordResponse(w http.ResponseWriter, r *http.Request, err error) {
	// Record Not Found generic error response
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(err.Error()))
}

func CreatedResponse(w http.ResponseWriter, r *http.Request, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payload)
}

func UpdatedResponse(w http.ResponseWriter, r *http.Request, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

func UnprocessableEntityResponse(w http.ResponseWriter, r *http.Request, err error, msg string) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte(msg))
	w.Write([]byte("\n"))
}

func UnavailableResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write([]byte(err.Error()))
}
