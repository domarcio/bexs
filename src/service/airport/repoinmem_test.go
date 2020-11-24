package airport

import (
	"context"
	"testing"
	"time"

	"github.com/domarcio/bexs/src/entity"
)

func TestGetRepo(t *testing.T) {
	repo := newRepoInmem()

	t.Run("successful", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
		defer cancel()

		repo.addToMemory("foo")

		airport, err := repo.Get(ctx, "foo")
		if err != nil {
			t.Errorf("unxpected error: %s", err.Error())
		}
		if airport == nil {
			t.Error("expected an airport")
		}
	})

	t.Run("timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
		defer cancel()
		time.Sleep(5 * time.Millisecond)

		airport, err := repo.Get(ctx, "foo")
		exp := entity.ErrTimeoutExceeded
		if err != exp {
			t.Errorf("expected error: %s", exp.Error())
		}
		if airport != nil {
			t.Error("unexpected airport found")
		}
	})
}
