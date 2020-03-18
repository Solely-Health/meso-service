package facilities

import (
	"fmt"

	"github.com/meso-org/meso/repository"
)

type service struct {
	facility  repository.FacilityRepository
	positions repository.PositionRepository
}

func (s *service) RegisterNewFacility(facilityName, email string) (repository.FacilityID, error) {
	if email == "" || facilityName == "" {
		return "", fmt.Errorf("in RegisterNewWorker, provided arguments are invalid")
	}

	facilityID := repository.GenerateFacilityID()

	facility := repository.NewFacility(facilityID, facilityName, email)

	if err := s.facility.Store(facility); err != nil {
		return "", err
	}

	return facility.FacilityID, nil
}

func (s *service) CreateJob(facilityID repository.FacilityID, positionID repository.PositionID, schedules []repository.Schedule, title, description string) (repository.Position, error) {
	newPosition := repository.Position{
		PositionID:  positionID,
		FacilityID:  facilityID,
		Description: description,
		Title:       title,
		Schedule:    schedules,
	}

	if err := s.positions.Store(&newPosition); err != nil {
		//TODO: dont return positionID lol
		return newPosition, fmt.Errorf("Unable to execute CreateJob: %v", err)
	}
	// TODO: Create an event that notifies workers that a job has been posted (if its in their location range)
	return newPosition, nil
}

func NewService(facilityRepo repository.FacilityRepository, positionsRepo repository.PositionRepository) Service {
	return &service{
		//TODO make plural lol
		facility:  facilityRepo,
		positions: positionsRepo,
	}
}
