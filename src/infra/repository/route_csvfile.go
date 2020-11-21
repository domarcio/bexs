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
func (repo *RouteCSVFile) Create(ctx context.Context, e *entity.Route) error {
	return nil
}

// ListFrom routes
func (repo *RouteCSVFile) ListFrom(ctx context.Context, from string) ([]*entity.Route, error) {
	return nil, nil
}
