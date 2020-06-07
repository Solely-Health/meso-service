package workers

import (
	"fmt"
	"strconv"

	"github.com/asaskevich/EventBus"
	"github.com/meso-org/meso/repository"
)

type Service interface {
	RegisterNewWorker(email, firstName, lastName, occupation, license string) (repository.WorkerID, error)
	FindWorkerByEmail(repository.Email) (*repository.Worker, error)
	FindWorkerByID(repository.WorkerID) (*repository.Worker, error)
	UpdateWorkerLocationPreference(workerID, latitude, longitude, mileRadius string) (*repository.Worker, error)
	FindAllWorkers() ([]*repository.Worker, error)
}

type service struct {
	workers  repository.WorkerRepository
	eventBus EventBus.Bus
}

func (s *service) RegisterNewWorker(email, firstName, lastName, occupation, license string) (repository.WorkerID, error) {
	// TODO Skills, range
	// email, first name, last name, password, licenses,
	if email == "" || firstName == "" || lastName == "" || occupation == "" {
		return "", fmt.Errorf("in RegisterNewWorker, provided arguments are invalid")
	}

	workerID := repository.GenerateWorkerID()

	parsedEmail := repository.Email(email)
	defaultLocation := repository.DefaultLocation()

	worker := repository.NewWorker(workerID, *defaultLocation, parsedEmail, firstName, lastName, occupation, license)
	if err := s.workers.Store(worker); err != nil {
		return "", err
	}
	// s.eventBus.Publish("workers:NewWorkerRegistered", worker)
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

func (s *service) UpdateWorkerLocationPreference(workerID, latitude, longitude, mileRadius string) (*repository.Worker, error) {
	var err error
	lat, err := strconv.ParseFloat(latitude, 8)
	fmt.Printf("%T, %v\n", lat, lat)
	if err != nil {
		return nil, err
	}

	long, err := strconv.ParseFloat(latitude, 8)
	fmt.Printf("%T, %v\n", lat, lat)
	if err != nil {
		return nil, err
	}

	parsedLocation := repository.NewLocation(lat, long)
	parsedWorkerID := repository.WorkerID(workerID)
	fmt.Println("HEYYYYYY: ", workerID, parsedWorkerID)

	if err = s.verifyWorkerIDCurrentlyExists(parsedWorkerID); err != nil {
		fmt.Printf("ERROR: %v", err)
		return nil, err
	}
	worker, err := s.workers.Update(parsedWorkerID, parsedLocation)
	if err != nil {
		err = fmt.Errorf("Unable to perform UpdateWorkerLocationPreference in service: %v", err)
		fmt.Printf("ERROR: %v", err)
		return nil, err
	}

	// TODO EVENT: fire off LocationUpdated event
	return worker, nil
}

func (s *service) verifyWorkerIDCurrentlyExists(id repository.WorkerID) error {
	_, err := s.workers.Find(id)

	if err != nil {
		return err
	}
	return nil
}

// NewService - pass this function a repository instance,
// and it will return a new service that has access to that repository
func NewService(workersRepo repository.WorkerRepository, eventBus EventBus.Bus) Service {
	return &service{
		eventBus: eventBus,
		workers:  workersRepo,
	}
}
