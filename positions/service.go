package positions

import "github.com/meso-org/meso/repository"

type service struct {
	positions repository.PositionRepository
	workers   repository.WorkerRepository
}

func (s *service) AddWorkerToSchedule(worker repository.WorkerID) {

}
