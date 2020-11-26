package cost

import (
	"fmt"
	"strings"

	"github.com/domarcio/bexs/src/entity"
	commonLog "github.com/domarcio/bexs/src/infra/log"
	"github.com/domarcio/bexs/src/service/connection"
)

// Service cost
type Service struct {
	connService connection.Servicer
	log         commonLog.Logger
}

// NewService create a new service to management costs
func NewService(connService connection.Servicer, log commonLog.Logger) *Service {
	return &Service{
		connService: connService,
		log:         log,
	}
}

// LowCost returns a string with better route on regardless number of connections
func (s *Service) LowCost(source *entity.Airport, target *entity.Airport) (string, error) {
	s.log.Info("Will to the check the best option (cost) for target %s and source %s", source.Code, target.Code)
	var (
		cc  = newTravelConnections(s.connService)
		err = cc.availableConnections(source, target, source, []string{}, 0, 0)
		c   = cc.connections
	)
	if err != nil {
		s.log.Error("Oops, an error occurred find the available connections: %s", err.Error())
		return "", err
	}

	if len(c) <= 0 {
		s.log.Warning("No available connections")
		return "", nil
	}

	var i informations
	var k float64
	for index, connections := range c {
		if connections.price < k || k == 0 {
			k = connections.price
			i = connections
		}

		s.log.Info("[%d] Route: %+v", index, connections.connections)
	}

	lowCostRoute := formatConnection(i.connections, i.price)
	s.log.Info("The low cost route is: %s", lowCostRoute)
	return lowCostRoute, nil
}

func formatConnection(s []string, price float64) string {
	return fmt.Sprintf("%s > $%.0f", strings.Join(s, " - "), price)
}
