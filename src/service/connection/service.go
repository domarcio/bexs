package connection

import (
	"context"
	"time"

	"github.com/domarcio/bexs/src/entity"
)

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
	connection, err := entity.NewConnection(source, target, price)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	alreadyExists, err := s.repo.Get(ctx, source, target)
	if err != nil {
		return nil, err
	}
	if alreadyExists != nil {
		return nil, entity.ErrConnectionAlreadyExists
	}

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Millisecond)
	defer cancel()

	err = s.repo.Create(ctx, connection)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

// FindConnections list all routes by source
func (s *Service) FindConnections(source *entity.Airport) ([]*entity.Connection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	return s.repo.ListBySource(ctx, source)
}
