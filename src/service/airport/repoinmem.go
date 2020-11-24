package airport

import (
	"context"
	"time"

	"github.com/domarcio/bexs/src/entity"
)

type repoinmem struct {
	m map[string]*entity.Airport
}

func newRepoInmem() *repoinmem {
	var m = make(map[string]*entity.Airport, 0)
	return &repoinmem{
		m: m,
	}
}

func (repo *repoinmem) addToMemory(code string) {
	_, ok := repo.m[code]
	if ok {
		return
	}

	newAirport := repo.m
	newAirport[code] = &entity.Airport{Code: code}

	repo.m = newAirport
}

func (repo *repoinmem) reset() {
	repo.m = make(map[string]*entity.Airport, 0)
}

func (repo *repoinmem) Get(ctx context.Context, code string) (*entity.Airport, error) {
	airport, ok := repo.m[code]

	select {
	case <-time.After(2 * time.Millisecond):
		break
	case <-ctx.Done():
		return nil, entity.ErrTimeoutExceeded
	}

	if !ok {
		return nil, nil
	}

	return airport, nil
}
