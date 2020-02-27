package workers

import (
	"testing"

	"github.com/meso-org/meso/repository"
)

type mockWorkersRepository struct {
	worker *repository.Worker
}

func (mockWorker *mockWorkersRepository) Store(worker *repository.Worker) error {
	mockWorker.worker = worker
	return nil
}
func (mockWorker *mockWorkersRepository) Find(x interface{}) (*repository.Worker, error) {
	return nil, nil
}
func (mockWorker *mockWorkersRepository) FindAll() ([]*repository.Worker, error) {
	return nil, nil
}
func TestRegisterNewWorker(t *testing.T) {
	// create mock variables
	var (
		email      = "mcnultymikkal@gmail.com"
		firstName  = "Mike"
		lastName   = "McNut"
		occupation = "Respitory Therapist"
		license    = "ASD123"
	)
	var workers mockWorkersRepository

	s := NewService(&workers)
	id, err := s.RegisterNewWorker(email, firstName, lastName, occupation, license)
	if err != nil {
		t.Fatal(err)
	}
	if id == "" {
		t.Fatal(id)
	}
}

// func TestFindWorkerByID(t *testing.T) {
// 	var (
// 		email      = "mcnultymikkal@gmail.com"
// 		firstName  = "Mike"
// 		lastName   = "McNut"
// 		occupation = "Respitory Therapist"
// 		license    = "ASD123"
// 	)
// 	var workers mockWorkersRepository

// 	s := NewService(&workers)
// 	id, err := s.RegisterNewWorker(email, firstName, lastName, occupation, license)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if id == "" {
// 		t.Fatal(id)
// 	}

// 	w, err := s.FindWorkerByID(id)
// 	fmt.Println(ww)
// 	if w == nil {
// 		t.Fatal(err)
// 	}
// }
