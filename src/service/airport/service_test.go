package airport

import (
	"testing"
)

func TestAirportExists(t *testing.T) {
	repo := newRepoInmem()
	service := NewService(repo)

	t.Run("false", func(t *testing.T) {
		exists, err := service.AirportExists("xxx")
		if err != nil {
			t.Errorf("unexpected error: %s", err.Error())
		}
		if exists {
			t.Errorf("it's not expected to find an airport")
		}
	})

	t.Run("true", func(t *testing.T) {
		repo.addToMemory("axy")
		exists, err := service.AirportExists("axy")
		if err != nil {
			t.Errorf("unexpected error: %s", err.Error())
		}
		if !exists {
			t.Errorf("expected to find an airport")
		}
	})
}
