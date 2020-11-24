package connection

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

// CreateConnection a new route on storage
func (s *Service) CreateConnection(source *entity.Airport, target *entity.Airport, price float64) (*entity.Connection, error) {
	return nil, nil
}

// FindConnections list all routes by source
func (s *Service) FindConnections(source *entity.Airport) ([]*entity.Connection, error) {
	return nil, nil
}
