package route

import (
	"context"

	"github.com/domarcio/bexs/src/entity"
)

// Repository interface
type Repository interface {
	Create(ctx context.Context, e *entity.Route) error
	ListFrom(ctx context.Context, from string) ([]*entity.Route, error)
}

// Servicer inteface
type Servicer interface {
	CreateRoute(from string, to string, price float64) (*entity.Route, error)
	FindRoutesByOrigin(from string) ([]*entity.Route, error)
}
