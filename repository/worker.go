package repository

import (
	"github.com/beevik/guid"
)

type WorkerID string
type Email string

// Location: TODO move this into its own package
type Location struct {
	Latitude  float64
	Longitude float64
}

func NewLocation(lat, long float64) *Location {
	return &Location{
		Latitude:  lat,
		Longitude: long,
	}
}

func DefaultLocation() *Location {
	return &Location{
		Latitude:  0,
		Longitude: 0,
	}
}

type WorkerRepository interface {
	Store(worker *Worker) error
	Find(x interface{}) (*Worker, error)
	FindAll() ([]*Worker, error)
	Update(workerID WorkerID, x interface{}) (*Worker, error)
}

// Domain object
type Worker struct {
	WorkerID  WorkerID
	Email     Email
	FirstName string
	LastName  string
	// TODO make this a constant enum
	Occupation string
	// TODO make a license a new defined type
	License string
	// TODO
	LocationPreference Location
}

// NewWorker - generate a new worker domain object with provided fields
func NewWorker(workerID WorkerID, location Location, email Email, firstName, lastName, occupation, license string) *Worker {

	return &Worker{
		WorkerID:           workerID,
		Email:              email,
		FirstName:          firstName,
		LastName:           lastName,
		Occupation:         occupation,
		License:            license,
		LocationPreference: location,
	}
}

// GenerateWorkerID - return a new WorkerID string
func GenerateWorkerID() WorkerID {
	return WorkerID(guid.NewString())
}
