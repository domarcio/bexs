package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/domarcio/bexs/src/entity"
)

type costService struct{}

func (s *costService) LowCost(source *entity.Airport, target *entity.Airport) (string, error) {
	if source.Code == "ERR" && target.Code == "ERR" {
		return "", errors.New("internal server error")
	}
	if source.Code == "GRU" && target.Code == "CDG" {
		return "GRU - CDG > $1", nil
	}

	return "", nil
}

func TestLow(t *testing.T) {
	handler := NewCostHandlers(&costService{})
	endpoint := "/api/cost"

	sm := http.NewServeMux()
	sm.HandleFunc(endpoint, handler.Low)

	ts := httptest.NewServer(sm)
	defer ts.Close()

	t.Run("invalid_parameters", func(t *testing.T) {
		res, err := http.Get(ts.URL + endpoint)
		if err != nil {
			t.Error("unexpected error")
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, res.StatusCode)
		}

		res, err = http.Get(ts.URL + endpoint + "?source=FOO")
		if err != nil {
			t.Error("unexpected error")
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, res.StatusCode)
		}

		res, err = http.Get(ts.URL + endpoint + "?target=BAR")
		if err != nil {
			t.Error("unexpected error")
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, res.StatusCode)
		}
	})

	t.Run("internal_server_error", func(t *testing.T) {
		res, err := http.Get(ts.URL + endpoint + "?source=ERR&target=ERR")
		if err != nil {
			t.Error("unexpected error")
		}
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, res.StatusCode)
		}
	})

	t.Run("not_found", func(t *testing.T) {
		res, err := http.Get(ts.URL + endpoint + "?source=GRU&target=CDG")
		if err != nil {
			t.Error("unexpected error")
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
		}
	})
}
