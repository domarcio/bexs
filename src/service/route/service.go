package route

import "github.com/domarcio/bexs/src/entity"

// Service route
type Service struct {
	repo Repository
}

// NewService create a new service to management routes
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// CreateRoute a new route on storage
func (s *Service) CreateRoute(from string, to string, price float64) (*entity.Route, error) {
	return nil, nil
}

// FindRoutesByOrigin list all routes by from/origin
func (s *Service) FindRoutesByOrigin(from string) ([]*entity.Route, error) {
	return nil, nil
}
