package handler

import "net/http"

type createAirport struct {
	Code string `json:"code"`
}

type listAirport struct {
	List []createAirport `json:"airports"`
}

// AirportHandlers struct handler
type AirportHandlers struct{}

// NewAirportHandlers create a new airport handlers
func NewAirportHandlers() *AirportHandlers {
	return &AirportHandlers{}
}

// Create handler
func (a *AirportHandlers) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
