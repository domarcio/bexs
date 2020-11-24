package airport

import (
	"context"

	"github.com/domarcio/bexs/src/entity"
)

// Repository interface
type Repository interface {
	Get(ctx context.Context, code string) (*entity.Airport, error)
}

// Servicer inteface
type Servicer interface {
	AirportExists(code string) (bool, error)
}
