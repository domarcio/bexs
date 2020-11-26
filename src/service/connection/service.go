package connection

import (
	"context"
	"time"

	"github.com/domarcio/bexs/src/entity"
	commonLog "github.com/domarcio/bexs/src/infra/log"
)

// Service route
type Service struct {
	repo Repository
	log  commonLog.Logger
}

// NewService create a new service to management routes
func NewService(repo Repository, log commonLog.Logger) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}

// CreateConnection a new route on storage
func (s *Service) CreateConnection(source *entity.Airport, target *entity.Airport, price float64) (*entity.Connection, error) {
	s.log.Info("Create connection received `%s` as source, `%s` as target and `%.0f` as price", source.Code, target.Code, price)

	connection, err := entity.NewConnection(source, target, price)
	if err != nil {
		s.log.Error("Oops, an error occurred on create a new connection: %s", err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	alreadyExists, err := s.repo.Get(ctx, source, target)
	if err != nil {
		s.log.Error("Oops, an error occurred on get the connection: %s", err.Error())
		return nil, err
	}
	if alreadyExists != nil {
		s.log.Error("Oops, the connection already exists")
		return nil, entity.ErrConnectionAlreadyExists
	}

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Millisecond)
	defer cancel()

	err = s.repo.Create(ctx, connection)
	if err != nil {
		s.log.Error("Oops, an error occurred on create the connection: %s", err.Error())
		return nil, err
	}

	s.log.Info("Connection created successfully")
	return connection, nil
}

// FindConnections list all routes by source
func (s *Service) FindConnections(source *entity.Airport) ([]*entity.Connection, error) {
	s.log.Info("Find connections by source `%s`", source.Code)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	list, err := s.repo.ListBySource(ctx, source)
	if err != nil {
		s.log.Error("Oops, an error occurred on find connections by source: %s", err.Error())
		return nil, err
	}

	s.log.Info("Found %d connections", len(list))
	return list, nil
}
