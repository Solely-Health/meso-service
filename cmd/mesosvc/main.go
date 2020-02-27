package main

import (
	"fmt"
	"net/http"

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
		workersRepo repo.WorkerRepository
	)

	// For development purposes we will just use in memory db (for now, can be configured)
	if inmemorydb {
		workersRepo = inmem.NewWorkerRepository()
	} else {
		// we can pick and choose what kind of db we want to use here
	}

	// Service Registration here
	var workersSVC workers.Service
	workersSVC = workers.NewService(workersRepo)

	srv := server.New(workersSVC)
	fmt.Println("bout to serve")
	http.ListenAndServe(":4040", srv)
}
