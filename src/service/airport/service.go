package airport

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
func (s *Service) AirportExists(code string) error {
	return nil
}
