package workers

import (
	"fmt"

	"github.com/meso-org/meso/repository"
)

type Service interface {
	RegisterNewWorker(email, firstName, lastName, occupation, license string) (repository.WorkerID, error)
	FindWorkerByEmail(repository.Email) (*repository.Worker, error)
	FindWorkerByID(repository.WorkerID) (*repository.Worker, error)
	FindAllWorkers() ([]*repository.Worker, error)
}

type service struct {
	workers repository.WorkerRepository
}

func (s *service) RegisterNewWorker(email, firstName, lastName, occupation, license string) (repository.WorkerID, error) {
	// TODO Skills, range
	// email, first name, last name, password, licenses,
	if email == "" || firstName == "" || lastName == "" || occupation == "" {
		return "", fmt.Errorf("in RegisterNewWorker, provided arguments are invalid")
	}

	workerID := repository.GenerateWorkerID()

	parsedEmail := repository.Email(email)

	worker := repository.NewWorker(workerID, parsedEmail, firstName, lastName, occupation, license)
	if err := s.workers.Store(worker); err != nil {
		return "", err
	}

	// we can trigger a "NewWorkerRegistered" to other services from here
	return worker.WorkerID, nil
}

func (s *service) FindWorkerByEmail(email repository.Email) (*repository.Worker, error) {
	w := repository.Worker{}
	if email == "" {
		return &w, fmt.Errorf("Bad request for FindWorkerByEmail, missing email")
	}

	worker, err := s.workers.Find(email)
	if err != nil {
		return nil, err
	}

	return worker, nil
}

func (s *service) FindWorkerByID(id repository.WorkerID) (*repository.Worker, error) {
	w := repository.Worker{}
	if id == "" {
		return &w, fmt.Errorf("Bad request for FindWorkerById, missing id")
	}

	worker, err := s.workers.Find(id)
	if err != nil {
		return nil, err
	}

	return worker, nil
}

// FindAllWorkers - return all workers, we can add ability to find all workers by their location, etc, etc
func (s *service) FindAllWorkers() ([]*repository.Worker, error) {
	workers, err := s.workers.FindAll()
	if err != nil {
		return nil, err
	}
	return workers, err
}

// NewService - pass this function a repository instance,
// and it will return a new service that has access to that repository
func NewService(workersRepo repository.WorkerRepository) Service {
	return &service{
		workers: workersRepo,
	}
}
