package cost

import (
	"github.com/domarcio/bexs/src/entity"
	"github.com/domarcio/bexs/src/service/connection"
)

type index uint

type informations struct {
	connections []string
	price       float64
}

type travel struct {
	connService connection.Servicer
	connections map[index]informations
}

func newTravelConnections(connService connection.Servicer) *travel {
	return &travel{
		connService: connService,
		connections: make(map[index]informations, 0),
	}
}

func (f *travel) availableConnections(source *entity.Airport, target *entity.Airport, origin *entity.Airport, connections []string, price float64, key index) error {
	connection, err := f.connService.FindConnections(source)
	if err != nil {
		return err
	}

	for _, c := range connection {
		if source.Code == origin.Code {
			price = 0
			connections = []string{origin.Code}
			key++
		}

		connections = append(connections, c.Target.Code)
		price += c.Price

		if c.Target.Code == target.Code {
			newConnections := f.connections
			newConnections[key] = informations{connections: connections, price: price}
			f.connections = newConnections
			connections = []string{}
			price = 0
			continue
		}

		if c.Target.Code != target.Code {
			f.availableConnections(c.Target, target, origin, connections, price, key)
		}
	}

	return nil
}
