package positions

import (
	"github.com/meso-org/meso/repository"
)

type Service interface {
	CreateNewPosition(repository.FacilityID, repository.StartDateTime, repository.EndDateTime) (repository.PositionID, error)
	FindPositionByID(repository.PositionID) (*repository.Position, error)
	FindAllPositions() ([]*repository.Position, error)
	AddWorkerToSchedule(worker repository.WorkerID)
}
