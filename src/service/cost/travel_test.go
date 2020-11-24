package cost

import (
	"fmt"
	"testing"

	"github.com/domarcio/bexs/src/entity"
)

func TestAvailableConnections(t *testing.T) {
	connService := newConnectionService()
	cc := newTravelConnections(connService)

	t.Run("source_not_found", func(t *testing.T) {
		cc.reset()
		cc.availableConnections(
			&entity.Airport{Code: "FOO"},
			&entity.Airport{Code: "CDG"},
			&entity.Airport{Code: "FOO"},
			[]string{},
			0,
			0,
		)
		c := cc.connections
		l := len(c)
		fmt.Println(c)
		if l > 0 {
			t.Errorf("unexpected found possible connections, got %d", l)
		}
	})

	t.Run("target_not_found", func(t *testing.T) {
		cc.reset()
		cc.availableConnections(
			&entity.Airport{Code: "GRU"},
			&entity.Airport{Code: "BAR"},
			&entity.Airport{Code: "GRU"},
			[]string{},
			0,
			0,
		)
		c := cc.connections
		l := len(c)
		fmt.Println(c)
		if l > 0 {
			t.Errorf("unexpected found possible connections, got %d", l)
		}
	})

	t.Run("successful_runtimeoptions", func(t *testing.T) {
		connService.CreateConnection(&entity.Airport{Code: "AAA"}, &entity.Airport{Code: "BBB"}, 1)
		connService.CreateConnection(&entity.Airport{Code: "BBB"}, &entity.Airport{Code: "CCC"}, 2)
		connService.CreateConnection(&entity.Airport{Code: "CCC"}, &entity.Airport{Code: "DDD"}, 3)
		connService.CreateConnection(&entity.Airport{Code: "DDD"}, &entity.Airport{Code: "EEE"}, 4)
		connService.CreateConnection(&entity.Airport{Code: "AAA"}, &entity.Airport{Code: "EEE"}, 20)

		expectedRoutes := []string{
			"AAA - BBB - CCC - DDD - EEE > $10",
			"AAA - EEE > $20",
		}

		for i := 0; i < 1000; i++ {
			cc.reset()
			cc.availableConnections(
				&entity.Airport{Code: "AAA"},
				&entity.Airport{Code: "EEE"},
				&entity.Airport{Code: "AAA"},
				[]string{},
				0,
				0,
			)
			c := cc.connections

			for _, connections := range c {
				r := formatConnection(connections.connections, connections.price)
				if !contains(r, expectedRoutes) {
					t.Fatalf("[%d] expected `%s` as result in list", i, r)
				}
			}
		}
	})
}

func contains(needle string, haystack []string) bool {
	for _, h := range haystack {
		if needle == h {
			return true
		}
	}
	return false
}
