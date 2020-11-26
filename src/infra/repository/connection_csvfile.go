package repository

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/domarcio/bexs/src/entity"
	"github.com/domarcio/bexs/src/infra/file"
)

// RouteCSVFile repository
type RouteCSVFile struct {
	writer file.Writer
	reader file.Reader
}

// NewRouteCSVFile create new repository
func NewRouteCSVFile(w file.Writer, r file.Reader) (*RouteCSVFile, error) {
	routeCSV := &RouteCSVFile{
		writer: w,
		reader: r,
	}
	return routeCSV, nil
}

// Create a new route
func (repo *RouteCSVFile) Create(ctx context.Context, e *entity.Connection) error {
	txt := fmt.Sprintf("%s,%s,%.0f\n", e.Source.Code, e.Target.Code, e.Price)
	return repo.writer.Append(txt)
}

// ListBySource routes
func (repo *RouteCSVFile) ListBySource(ctx context.Context, source *entity.Airport) ([]*entity.Connection, error) {
	list := make([]*entity.Connection, 0)

	select {
	case <-time.After(4 * time.Millisecond):
		break
	case <-ctx.Done():
		return nil, entity.ErrTimeoutExceeded
	}

	defer repo.reader.Rewind()

	for repo.reader.Valid() {
		repo.reader.Next()

		if repo.reader.Error() == io.EOF {
			break
		}
		if err := repo.reader.Error(); err != nil {
			return nil, err
		}

		line := repo.reader.Current()
		if len(line) < 3 {
			continue
		}

		parsePrice, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			return nil, err
		}

		connection := &entity.Connection{
			Source: &entity.Airport{Code: line[0]},
			Target: &entity.Airport{Code: line[1]},
			Price:  parsePrice,
		}
		if connection.Source.Code == source.Code {
			list = append(list, connection)
		}
	}

	return list, nil
}

// Get returns a single result
func (repo *RouteCSVFile) Get(ctx context.Context, source *entity.Airport, target *entity.Airport) (*entity.Connection, error) {
	select {
	case <-time.After(4 * time.Millisecond):
		break
	case <-ctx.Done():
		return nil, entity.ErrTimeoutExceeded
	}

	defer repo.reader.Rewind()

	for repo.reader.Valid() {
		repo.reader.Next()

		if repo.reader.Error() == io.EOF {
			break
		}
		if err := repo.reader.Error(); err != nil {
			return nil, err
		}

		line := repo.reader.Current()
		if len(line) < 3 {
			continue
		}

		parsePrice, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			return nil, err
		}

		connection := &entity.Connection{
			Source: &entity.Airport{Code: line[0]},
			Target: &entity.Airport{Code: line[1]},
			Price:  parsePrice,
		}

		if connection.Source.Code == source.Code && connection.Target.Code == target.Code {
			return connection, nil
		}
	}

	return nil, nil
}
