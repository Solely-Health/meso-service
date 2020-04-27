package main

import (
	"fmt"
	"net/http"

	"github.com/meso-org/meso/facilities"
	inmem "github.com/meso-org/meso/inmemorydb"
	repo "github.com/meso-org/meso/repository"
	server "github.com/meso-org/meso/server"
	workers "github.com/meso-org/meso/workers"
)

func main() {
	var (
		inmemorydb = true
	)

	// Repository Registration here
	var (
		workersRepo    repo.WorkerRepository
		facilitiesRepo repo.FacilityRepository
	)

	// For development purposes we will just use in memory db (for now, can be configured)
	if inmemorydb {
		workersRepo = inmem.NewWorkerRepository()
		facilitiesRepo = inmem.NewFacilityRepository()
	} else {
		// we can pick and choose what kind of db we want to use here
	}

	// Service Registration here
	var workersSVC workers.Service
	workersSVC = workers.NewService(workersRepo)
	var facilitySVC facilities.Service
	facilitySVC = facilities.NewService(facilitiesRepo)

	srv := server.New(workersSVC, facilitySVC)
	fmt.Println("bout to serve")
	http.ListenAndServe(":4040", srv)
}
