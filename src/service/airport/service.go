package airport

import (
	"context"
	"time"
)

// Service route
type Service struct {
	repo Repository
}

// NewService create a new service to management airports
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// AirportExists check if an airport already exists by code
func (s *Service) AirportExists(code string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Millisecond)
	defer cancel()

	airport, err := s.repo.Get(ctx, code)
	if err != nil {
		return false, err
	}
	if airport == nil {
		return false, nil
	}

	return true, nil
}
