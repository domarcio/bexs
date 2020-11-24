package connection

import (
	"context"

	"github.com/domarcio/bexs/src/entity"
)

// Repository interface
type Repository interface {
	Create(ctx context.Context, e *entity.Connection) error
	ListBySource(ctx context.Context, source *entity.Airport) ([]*entity.Connection, error)
	Get(ctx context.Context, source *entity.Airport, target *entity.Airport) (*entity.Connection, error)
}

// Servicer inteface
type Servicer interface {
	CreateConnection(source *entity.Airport, target *entity.Airport, price float64) (*entity.Connection, error)
	FindConnections(source *entity.Airport) ([]*entity.Connection, error)
}
