package positions

import (
	"github.com/meso-org/meso/repository"
)

type Service interface {
	AddWorkerToSchedule(worker repository.WorkerID)
}
