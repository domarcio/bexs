package repository

import (
	"context"

	"github.com/domarcio/bexs/src/entity"
)

// RouteCSVFile repository
type RouteCSVFile struct {
	filename string
}

// NewRouteCSVFile create new repository
func NewRouteCSVFile(filename string) *RouteCSVFile {
	return &RouteCSVFile{
		filename: filename,
	}
}

// Create a new route
func (repo *RouteCSVFile) Create(ctx context.Context, e *entity.Connection) error {
	return nil
}

// ListBySource routes
func (repo *RouteCSVFile) ListBySource(ctx context.Context, source *entity.Airport) ([]*entity.Connection, error) {
	return nil, nil
}

// Get returns a single result
func (repo *RouteCSVFile) Get(ctx context.Context, source *entity.Airport, target *entity.Airport) (*entity.Connection, error) {
	return nil, nil
}
