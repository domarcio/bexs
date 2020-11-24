package connection

import (
	"testing"

	"github.com/domarcio/bexs/src/entity"
)

func TestCreateConnection(t *testing.T) {
	service := NewService(newRepoInmem())

	t.Run("error", func(t *testing.T) {
		_, err := service.CreateConnection(&entity.Airport{Code: "FOO"}, &entity.Airport{Code: "BAR"}, 0)
		if err == nil {
			t.Error("expected found an error")
		}
	})

	t.Run("successful", func(t *testing.T) {
		connection, err := service.CreateConnection(&entity.Airport{Code: "FOO"}, &entity.Airport{Code: "BAR"}, 1)
		if err != nil {
			t.Error("unexpected error")
		}
		if connection.Source.Code != "FOO" {
			t.Errorf("expected FOO, got %s", connection.Source)
		}
		if connection.Target.Code != "BAR" {
			t.Errorf("expected BAR, got %s", connection.Source)
		}
		if connection.Price != 1 {
			t.Errorf("expected 1, got %.0f", connection.Price)
		}
	})

	t.Run("already_exists", func(t *testing.T) {
		_, err := service.CreateConnection(&entity.Airport{Code: "FOO"}, &entity.Airport{Code: "BAR"}, 1)
		exp := entity.ErrConnectionAlreadyExists
		if err == nil {
			t.Error("expected found an error")
		}
		if err != exp {
			t.Errorf("expected error %s, got %s", exp.Error(), err.Error())
		}
	})
}

func TestListBySource(t *testing.T) {
	service := NewService(newRepoInmem())

	t.Run("successful", func(t *testing.T) {
		service.CreateConnection(&entity.Airport{Code: "FOO"}, &entity.Airport{Code: "BAR"}, 1)
		service.CreateConnection(&entity.Airport{Code: "BAR"}, &entity.Airport{Code: "FOO"}, 1)

		list, err := service.FindConnections(&entity.Airport{Code: "BAR"})
		if err != nil {
			t.Errorf("unxpected error: %s", err.Error())
		}
		if len(list) != 1 {
			t.Error("expected 1 row")
		}
	})

	t.Run("not_found", func(t *testing.T) {
		list, err := service.FindConnections(&entity.Airport{Code: "XXX"})
		if err != nil {
			t.Errorf("unxpected error: %s", err.Error())
		}
		if len(list) > 0 {
			t.Error("expected 0 row")
		}
	})
}
