package utils

import (
	"fmt"
	"net/http"
)

// generic function to send an HTTP Response with payload
func PayloadResponse(w http.ResponseWriter, r *http.Request, payload *[]byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(*payload)
}

func NoRecordResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(err.Error()))
}

func CreatedResponse(w http.ResponseWriter, r *http.Request, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(payload)
	w.Write(payload.([]byte))
}

func UpdatedResponse(w http.ResponseWriter, r *http.Request, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(payload)
	w.Write(payload.([]byte))
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

func UnprocessableEntityResponse(w http.ResponseWriter, r *http.Request, err error, msg string) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte(msg))
}

func UnavailableResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write([]byte(err.Error()))
}

func MalformedIDResponse(w http.ResponseWriter, r *http.Request, t, id string) {
	msg := fmt.Sprintf("invalid %s id param: %s", t, id)
	ErrorResponse(w, r, http.StatusBadRequest, msg)
}

// migrate all error messages here
// should `msg` type be any?
func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, msg string) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}
