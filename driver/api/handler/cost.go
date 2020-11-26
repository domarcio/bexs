package handler

import (
	"net/http"

	"github.com/domarcio/bexs/src/entity"
	"github.com/domarcio/bexs/src/service/cost"
)

// CostHandlers struct handler
type CostHandlers struct {
	service cost.Servicer
}

type lowCostResponse struct {
	Route string `json:"route"`
}

// NewCostHandlers create a new airport handlers
func NewCostHandlers(service cost.Servicer) *CostHandlers {
	return &CostHandlers{
		service: service,
	}
}

// Low handler
func (c *CostHandlers) Low(w http.ResponseWriter, r *http.Request) {
	var (
		query  = r.URL.Query()
		source = query.Get("source")
		target = query.Get("target")
	)

	// Validate airport format
	airSource, err := entity.NewAirport(source)
	if err != nil {
		messageErr := structToString(responseError{
			Message: "Invalid source! " + err.Error(),
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(messageErr))
		return
	}

	airTarget, err := entity.NewAirport(target)
	if err != nil {
		messageErr := structToString(responseError{
			Message: "Invalid target! " + err.Error(),
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(messageErr))
		return
	}

	route, err := c.service.LowCost(airSource, airTarget)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	statusCode := http.StatusOK
	res := lowCostResponse{
		Route: route,
	}

	if route == "" {
		res.Route = "Route not found"
		statusCode = http.StatusNotFound
	}

	message := structToString(res)

	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
