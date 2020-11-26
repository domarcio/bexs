package handler

import (
	"encoding/json"
	"net/http"

	"github.com/domarcio/bexs/src/entity"
	"github.com/domarcio/bexs/src/service/connection"
)

// ConnectionHandlers struct handler
type ConnectionHandlers struct {
	service connection.Servicer
}

type createConnectionRequest struct {
	Source string
	Target string
	Price  float64
}

// NewConnectionHandlers create a new airport handlers
func NewConnectionHandlers(service connection.Servicer) *ConnectionHandlers {
	return &ConnectionHandlers{
		service: service,
	}
}

// Create handler
func (c *ConnectionHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var payload createConnectionRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var messageErr string
	source, err := entity.NewAirport(payload.Source)
	if err != nil {
		messageErr = structToString(responseError{
			Message: "Invalid payload! " + err.Error(),
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(messageErr))
		return
	}

	target, err := entity.NewAirport(payload.Target)
	if err != nil {
		messageErr = structToString(responseError{
			Message: "Invalid payload! " + err.Error(),
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(messageErr))
		return
	}

	_, err = c.service.CreateConnection(source, target, payload.Price)
	if err != nil {
		switch err {
		case entity.ErrMissingSourceOrTarget:
		case entity.ErrInvalidPrice:
			messageErr = structToString(responseError{
				Message: "Invalid payload! " + err.Error(),
			})
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(messageErr))
			return
		case entity.ErrConnectionAlreadyExists:
			messageErr = structToString(responseError{
				Message: "Invalid payload! " + err.Error(),
			})
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(messageErr))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
