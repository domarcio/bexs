package connection

import (
	"context"
	"testing"
	"time"

	"github.com/domarcio/bexs/src/entity"
)

var (
	dummySourceAirport = &entity.Airport{Code: "foo"}
	dummyTargetAirport = &entity.Airport{Code: "bar"}
)

func TestCreateRepo(t *testing.T) {
	repo := newRepoInmem()
	e := &entity.Connection{
		Source: dummySourceAirport,
		Target: dummyTargetAirport,
		Price:  10.50,
	}

	t.Run("successful", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
		defer cancel()

		err := repo.Create(ctx, e)
		if err != nil {
			t.Errorf("unxpected error: %s", err.Error())
		}
	})

	t.Run("timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
		defer cancel()
		time.Sleep(5 * time.Millisecond)

		err := repo.Create(ctx, e)
		exp := entity.ErrTimeoutExceeded
		if err != exp {
			t.Errorf("expected error: %s", exp.Error())
		}
	})
}

func TestListFromRepo(t *testing.T) {
	repo := newRepoInmem()

	t.Run("successful", func(t *testing.T) {
		repo.Create(context.Background(), &entity.Connection{
			Source: dummySourceAirport,
			Target: dummyTargetAirport,
			Price:  10.50,
		})
		repo.Create(context.Background(), &entity.Connection{
			Source: dummyTargetAirport,
			Target: &entity.Airport{Code: "xpto"},
			Price:  5,
		})
		repo.Create(context.Background(), &entity.Connection{
			Source: dummySourceAirport,
			Target: &entity.Airport{Code: "xpto"},
			Price:  5,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
		defer cancel()

		list, err := repo.ListBySource(ctx, dummySourceAirport)
		if err != nil {
			t.Errorf("unxpected error: %s", err.Error())
		}
		if len(list) != 2 {
			t.Error("expected 2 row")
		}
	})

	t.Run("timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
		defer cancel()
		time.Sleep(5 * time.Millisecond)

		_, err := repo.ListBySource(ctx, dummySourceAirport)
		exp := entity.ErrTimeoutExceeded
		if err != exp {
			t.Errorf("expected error: %s", exp.Error())
		}
	})
}
