package cost

import (
	"testing"

	"github.com/domarcio/bexs/src/entity"
	commonLog "github.com/domarcio/bexs/src/infra/log"
	"github.com/domarcio/bexs/src/service/connection"
)

var availableConnections = []*entity.Connection{
	{Source: &entity.Airport{Code: "GRU"}, Target: &entity.Airport{Code: "BRC"}, Price: 10},
	{Source: &entity.Airport{Code: "BRC"}, Target: &entity.Airport{Code: "SCL"}, Price: 5},
	{Source: &entity.Airport{Code: "GRU"}, Target: &entity.Airport{Code: "CDG"}, Price: 75},
	{Source: &entity.Airport{Code: "GRU"}, Target: &entity.Airport{Code: "SCL"}, Price: 20},
	{Source: &entity.Airport{Code: "GRU"}, Target: &entity.Airport{Code: "ORL"}, Price: 56},
	{Source: &entity.Airport{Code: "ORL"}, Target: &entity.Airport{Code: "CDG"}, Price: 5},
	{Source: &entity.Airport{Code: "SCL"}, Target: &entity.Airport{Code: "ORL"}, Price: 20},
}

type connectionService struct {
	availableConnections []*entity.Connection
}

func (c *connectionService) CreateConnection(source *entity.Airport, target *entity.Airport, price float64) (*entity.Connection, error) {
	connection, err := entity.NewConnection(source, target, price)
	if err != nil {
		return nil, err
	}

	c.availableConnections = append(c.availableConnections, connection)
	return connection, nil
}
func (c *connectionService) FindConnections(source *entity.Airport) ([]*entity.Connection, error) {
	var result []*entity.Connection

	for _, conn := range c.availableConnections {
		if conn.Source.Code == source.Code {
			result = append(result, conn)
		}
	}

	return result, nil
}

func newConnectionService() connection.Servicer {
	return &connectionService{
		availableConnections: availableConnections,
	}
}

func TestLowCost(t *testing.T) {
	s := NewService(newConnectionService(), commonLog.NewLogprint())

	t.Run("successful_gru_cdg", func(t *testing.T) {
		result, err := s.LowCost(&entity.Airport{Code: "GRU"}, &entity.Airport{Code: "CDG"})
		exp := "GRU - BRC - SCL - ORL - CDG > $40"

		if result == "" {
			t.Error("expected a non-empty result")
		}
		if err != nil {
			t.Error("unexpected error")
		}
		if result != exp {
			t.Errorf("expected a result as: %s", exp)
		}
	})

	t.Run("successful_orl_cdg", func(t *testing.T) {
		result, err := s.LowCost(&entity.Airport{Code: "ORL"}, &entity.Airport{Code: "CDG"})
		exp := "ORL - CDG > $5"

		if result == "" {
			t.Error("expected a non-empty result")
		}
		if err != nil {
			t.Error("unexpected error")
		}
		if result != exp {
			t.Errorf("expected a result as: %s", exp)
		}
	})

	t.Run("does_not_exists", func(t *testing.T) {
		result, err := s.LowCost(&entity.Airport{Code: "FOO"}, &entity.Airport{Code: "BAR"})
		exp := ""

		if err != nil {
			t.Error("unexpected error")
		}
		if result != exp {
			t.Errorf("expected a result as: %s", exp)
		}
	})
}
