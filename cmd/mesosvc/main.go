package main

import (
	"fmt"
	"net/http"

	"github.com/asaskevich/EventBus"
	"github.com/meso-org/meso/facilities"
	inmem "github.com/meso-org/meso/inmemorydb"
	"github.com/meso-org/meso/positions"
	repo "github.com/meso-org/meso/repository"
	server "github.com/meso-org/meso/server"
	workers "github.com/meso-org/meso/workers"
)

func main() {
	var (
		inmemorydb = true
		bus        = EventBus.New()
	)

	// Repository Registration here
	var (
		workersRepo    repo.WorkerRepository
		facilitiesRepo repo.FacilityRepository
		positionsRepo  repo.PositionRepository
	)

	// For development purposes we will just use in memory db (for now, can be configured)
	if inmemorydb {
		workersRepo = inmem.NewWorkerRepository()
		facilitiesRepo = inmem.NewFacilityRepository()
		positionsRepo = inmem.NewPositionRepository()
	} else {
		// we can pick and choose what kind of db we want to use here
	}

	// Service Registration here
	var workersSVC workers.Service
	workersSVC = workers.NewService(workersRepo, bus)
	var facilitySVC facilities.Service
	facilitySVC = facilities.NewService(facilitiesRepo, bus)
	var positionSVC positions.Service
	positionSVC = positions.NewService(positionsRepo, workersRepo)

	srv := server.New(workersSVC, facilitySVC, positionSVC)
	fmt.Println("bout to serve")
	http.ListenAndServe(":4040", srv)
}
