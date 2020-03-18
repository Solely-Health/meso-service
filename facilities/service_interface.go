package facilities

import "github.com/meso-org/meso/repository"

type Service interface {
	RegisterNewFacility(facilityName, email string) (repository.FacilityID, error)
	CreateJob(facilityID repository.FacilityID, positionID repository.PositionID, schedules []repository.Schedule, title, description string) (repository.Position, error)
}
