package connection

import (
	"context"
	"time"

	"github.com/domarcio/bexs/src/entity"
)

type repoinmem struct {
	m []*entity.Connection
}

func newRepoInmem() *repoinmem {
	var m = make([]*entity.Connection, 0)
	return &repoinmem{
		m: m,
	}
}

func (repo *repoinmem) Create(ctx context.Context, e *entity.Connection) error {
	select {
	case <-time.After(2 * time.Millisecond):
		break
	case <-ctx.Done():
		return entity.ErrTimeoutExceeded
	}

	repo.m = append(repo.m, e)

	return nil
}

func (repo *repoinmem) ListBySource(ctx context.Context, source *entity.Airport) ([]*entity.Connection, error) {
	list := make([]*entity.Connection, 0)

	select {
	case <-time.After(4 * time.Millisecond):
		break
	case <-ctx.Done():
		return nil, entity.ErrTimeoutExceeded
	}

	for _, e := range repo.m {
		if e.Source == source {
			list = append(list, e)
		}
	}

	return list, nil
}
