package handler

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/domarcio/bexs/src/entity"
	"github.com/domarcio/bexs/src/service/connection"
)

type repoService struct{}

func (*repoService) Create(ctx context.Context, e *entity.Connection) error {
	return nil
}
func (*repoService) ListBySource(ctx context.Context, source *entity.Airport) ([]*entity.Connection, error) {
	return nil, nil
}
func (*repoService) Get(ctx context.Context, source *entity.Airport, target *entity.Airport) (*entity.Connection, error) {
	return nil, nil
}

func TestCreate(t *testing.T) {
	handler := NewConnectionHandlers(connection.NewService(&repoService{}))
	endpoint := "/api/connection"

	sm := http.NewServeMux()
	sm.HandleFunc(endpoint, handler.Create)

	ts := httptest.NewServer(sm)
	defer ts.Close()

	t.Run("invalid_parameters", func(t *testing.T) {
		var reqBody *bytes.Buffer

		reqBody = bytes.NewBufferString("{}")
		res, err := http.Post(ts.URL+endpoint, "application/json", reqBody)
		if err != nil {
			t.Error("unexpected error")
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, res.StatusCode)
		}

		reqBody = bytes.NewBufferString(`{"source": "FOO", "price": 1}`)
		res, err = http.Post(ts.URL+endpoint, "application/json", reqBody)
		if err != nil {
			t.Error("unexpected error")
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, res.StatusCode)
		}

		reqBody = bytes.NewBufferString(`{"target": "FOO", "price": 1}`)
		res, err = http.Post(ts.URL+endpoint, "application/json", reqBody)
		if err != nil {
			t.Error("unexpected error")
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, res.StatusCode)
		}

		reqBody = bytes.NewBufferString(`{"source": "BAR", "target": "FOO"}`)
		res, err = http.Post(ts.URL+endpoint, "application/json", reqBody)
		if err != nil {
			t.Error("unexpected error")
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, res.StatusCode)
		}
	})
}
