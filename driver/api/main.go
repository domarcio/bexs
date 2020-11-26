package main

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/domarcio/bexs/config"
	"github.com/domarcio/bexs/driver/api/handler"
	"github.com/domarcio/bexs/src/infra/file"
	"github.com/domarcio/bexs/src/infra/repository"
	"github.com/domarcio/bexs/src/service/connection"
	"github.com/domarcio/bexs/src/service/cost"
)

func main() {
	log := config.LogService
	log.Info("Running api interface on `%s` environment", config.Env)

	write, read, err := file.NewCSVManager(config.RouteStorageFilePath)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	defer func() {
		write.CloseFile()
		read.CloseFile()
	}()

	repo, err := repository.NewRouteCSVFile(write, read)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	connService := connection.NewService(repo, log)
	costService := cost.NewService(connService, log)

	// Waiting for CTRL+C
	sg := make(chan os.Signal, 1)
	signal.Notify(sg, os.Interrupt)

	api := apiHandler{
		connectionService: connService,
		costService:       costService,
	}

	sm := http.NewServeMux()
	sm.Handle("/api/", api)

	go func() {
		http.ListenAndServe(":7007", sm)
	}()

	<-sg
	log.Info("Finish API")
}

type apiHandler struct {
	connectionService connection.Servicer
	costService       cost.Servicer
}

func (a apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resource := r.URL.Path[5:]

	switch resource {
	case "connection":
		setupConnection(a.connectionService, w, r)
		return
	case "cost":
		setupCost(a.costService, w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func setupConnection(service connection.Servicer, w http.ResponseWriter, r *http.Request) {
	connection := handler.NewConnectionHandlers(service)

	switch r.Method {
	case http.MethodPost:
		connection.Create(w, r)
		break
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func setupCost(service cost.Servicer, w http.ResponseWriter, r *http.Request) {
	airport := handler.NewCostHandlers(service)

	switch r.Method {
	case http.MethodGet:
		airport.Low(w, r)
		break
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
