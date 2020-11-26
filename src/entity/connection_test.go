package entity

import (
	"testing"
)

func TestNewConnection(t *testing.T) {
	t.Run("mising", func(t *testing.T) {
		_, err := NewConnection(nil, nil, 10)
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
	t.Run("price_wrong", func(t *testing.T) {
		_, err := NewConnection(&Airport{Code: "FOO"}, &Airport{Code: "BAR"}, 0)
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
	t.Run("same", func(t *testing.T) {
		_, err := NewConnection(&Airport{Code: "FOO"}, &Airport{Code: "FOO"}, 1)
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
	t.Run("successful", func(t *testing.T) {
		c, err := NewConnection(&Airport{Code: "FOO"}, &Airport{Code: "BAR"}, 1)
		if err != nil {
			t.Errorf("unexpected error")
		}
		if c == nil {
			t.Errorf("expected a connection struct, got nil")
		}
	})
}
