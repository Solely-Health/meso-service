package workers

import (
	"fmt"

	"github.com/meso-org/meso/repository"
)

func NewWorkerRegistered(worker repository.Worker) {
	fmt.Printf("Worker: %v", worker)
}
