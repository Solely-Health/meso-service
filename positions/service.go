package positions

import "github.com/meso-org/meso/repository"

type service struct {
	positions repository.PositionRepository
	workers   repository.WorkerRepository
}

func (s *service) AddWorkerToSchedule(worker repository.WorkerID) {

}

func (s *service) CreateNewPosition(facilityID repository.FacilityID,
	startDateTime repository.StartDateTime,
	endDateTime repository.EndDateTime,
	description, title string) (positionID repository.PositionID, err error) {
	newPosition := repository.Position{}

	positionID = repository.GeneratePositionID()

	newPosition.PositionID = positionID
	newPosition.StartDateTime = startDateTime
	newPosition.EndDateTime = endDateTime
	newPosition.FacilityID = facilityID
	newPosition.Description = description
	newPosition.Title = title

	err = s.positions.Store(&newPosition)
	if err != nil {
		return newPosition.PositionID, err
	}
	return newPosition.PositionID, nil
}

func (s *service) FindAllPositions() (positions []*repository.Position, err error) {
	positions, err = s.positions.FindAll()
	if err != nil {
		return nil, err
	}
	return positions, err
}

func NewService(positionRepo repository.PositionRepository, workersRepo repository.WorkerRepository) Service {
	return &service{
		positions: positionRepo,
		workers:   workersRepo,
	}
}
