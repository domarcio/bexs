package entity

import (
	"testing"
)

func TestNewAirport(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		_, err := NewAirport("")
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
	t.Run("greater_than_3", func(t *testing.T) {
		_, err := NewAirport("ABC1")
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
	t.Run("less_than_3", func(t *testing.T) {
		_, err := NewAirport("AB")
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
	t.Run("uppercase", func(t *testing.T) {
		_, err := NewAirport("Abc")
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
	t.Run("number", func(t *testing.T) {
		_, err := NewAirport("123")
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
	t.Run("successful", func(t *testing.T) {
		a, err := NewAirport("ABC")
		if err != nil {
			t.Errorf("unexpected error")
		}
		if a == nil {
			t.Errorf("expected a airport struct, got ni")
		}
		if a.Code != "ABC" {
			t.Errorf("expected ABC code, got %s", a.Code)
		}
	})
}
