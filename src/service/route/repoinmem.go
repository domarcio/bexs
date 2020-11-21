package route

import (
	"context"
	"time"

	"github.com/domarcio/bexs/src/entity"
)

type repoinmem struct {
	m []*entity.Route
}

func newRepoInmem() *repoinmem {
	var m = make([]*entity.Route, 0)
	return &repoinmem{
		m: m,
	}
}

func (repo *repoinmem) Create(ctx context.Context, e *entity.Route) error {
	//entitystr := fmt.Sprintf("%s,%s,%s", e.From, e.To, strconv.FormatFloat(e.Price, 'f', 6, 64))

	select {
	case <-time.After(2 * time.Millisecond):
		break
	case <-ctx.Done():
		return entity.ErrTimeoutExceeded
	}

	repo.m = append(repo.m, e)

	return nil
}

func (repo *repoinmem) ListFrom(ctx context.Context, from string) ([]*entity.Route, error) {
	list := make([]*entity.Route, 0)

	select {
	case <-time.After(4 * time.Millisecond):
		break
	case <-ctx.Done():
		return nil, entity.ErrTimeoutExceeded
	}

	for _, e := range repo.m {
		if e.From == from {
			list = append(list, e)
		}
	}

	return list, nil
}
