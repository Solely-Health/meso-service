package facilities

import "github.com/meso-org/meso/repository"

type Service interface {
	RegisterNewFacility() (repository.FacilityID, error)
	FindFacilityByID(repository.FacilityID) (*repository.Facility, error)
	FindAllFacilities() ([]*repository.Worker, error)
}
