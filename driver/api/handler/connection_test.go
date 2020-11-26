package handler

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/domarcio/bexs/src/entity"
	commonLog "github.com/domarcio/bexs/src/infra/log"
	"github.com/domarcio/bexs/src/service/connection"
)

type repoService struct{}

func (*repoService) Create(ctx context.Context, e *entity.Connection) error {
	if e.Source.Code == "AAA" && e.Target.Code == "BBB" {
		return entity.ErrTimeoutExceeded
	}

	return nil
}
func (*repoService) ListBySource(ctx context.Context, source *entity.Airport) ([]*entity.Connection, error) {
	return nil, nil
}
func (*repoService) Get(ctx context.Context, source *entity.Airport, target *entity.Airport) (*entity.Connection, error) {
	if source.Code == "BAR" && target.Code == "FOO" {
		return &entity.Connection{Source: source, Target: target, Price: 1}, nil
	}

	return nil, nil
}

func TestCreate(t *testing.T) {
	handler := NewConnectionHandlers(connection.NewService(&repoService{}, commonLog.NewLogprint()))
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

		reqBody = bytes.NewBufferString(`{"source": "123", "target": "456", "price": 1}`)
		res, err = http.Post(ts.URL+endpoint, "application/json", reqBody)
		if err != nil {
			t.Error("unexpected error")
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, res.StatusCode)
		}
	})

	t.Run("already_exists", func(t *testing.T) {
		var reqBody *bytes.Buffer

		reqBody = bytes.NewBufferString(`{"source": "BAR", "target": "FOO", "price": 1}`)
		res, err := http.Post(ts.URL+endpoint, "application/json", reqBody)
		if err != nil {
			t.Error("unexpected error")
		}
		if res.StatusCode != http.StatusConflict {
			t.Errorf("expected status code %d, got %d", http.StatusConflict, res.StatusCode)
		}
	})

	t.Run("internal_server_error", func(t *testing.T) {
		var reqBody *bytes.Buffer

		reqBody = bytes.NewBufferString(`{"source": "AAA", "target": "BBB", "price": 1}`)
		res, err := http.Post(ts.URL+endpoint, "application/json", reqBody)
		if err != nil {
			t.Error("unexpected error")
		}
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, res.StatusCode)
		}
	})

	t.Run("successful", func(t *testing.T) {
		var reqBody *bytes.Buffer

		reqBody = bytes.NewBufferString(`{"source": "BBB", "target": "CCC", "price": 1}`)
		res, err := http.Post(ts.URL+endpoint, "application/json", reqBody)
		if err != nil {
			t.Error("unexpected error")
		}
		if res.StatusCode != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, res.StatusCode)
		}
	})
}
