package cost

import "github.com/domarcio/bexs/src/entity"

// Servicer inteface
type Servicer interface {
	LowCost(source *entity.Airport, target *entity.Airport) (string, error)
}
