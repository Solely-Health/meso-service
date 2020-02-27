package repository

import (
	"github.com/beevik/guid"
)

type WorkerID string
type Email string

type WorkerRepository interface {
	Store(worker *Worker) error
	Find(x interface{}) (*Worker, error)
	FindAll() ([]*Worker, error)
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
}

// NewWorker - generate a new worker domain object with provided fields
func NewWorker(workerID WorkerID, email Email, firstName, lastName, occupation, license string) *Worker {

	return &Worker{
		WorkerID:   workerID,
		Email:      email,
		FirstName:  firstName,
		LastName:   lastName,
		Occupation: occupation,
		License:    license,
	}
}

// GenerateWorkerID - return a new WorkerID string
func GenerateWorkerID() WorkerID {
	return WorkerID(guid.NewString())
}
