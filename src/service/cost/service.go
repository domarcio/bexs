package cost

import (
	"fmt"
	"strings"

	"github.com/domarcio/bexs/src/entity"
	"github.com/domarcio/bexs/src/service/connection"
)

// Service cost
type Service struct {
	connService connection.Servicer
}

// NewService create a new service to management costs
func NewService(connService connection.Servicer) *Service {
	return &Service{
		connService: connService,
	}
}

// LowCost returns a string with better route on regardless number of connections
func (s *Service) LowCost(source *entity.Airport, target *entity.Airport) (string, error) {
	var (
		cc  = newTravelConnections(s.connService)
		err = cc.availableConnections(source, target, source, []string{}, 0, 0)
		c   = cc.connections
	)
	if err != nil {
		return "", err
	}

	if len(c) <= 0 {
		return "", nil
	}

	var i informations
	var k float64
	for _, connections := range c {
		if connections.price < k || k == 0 {
			k = connections.price
			i = connections
		}
	}

	return formatConnection(i.connections, i.price), nil
}

func formatConnection(s []string, price float64) string {
	return fmt.Sprintf("%s > $%.0f", strings.Join(s, " - "), price)
}
