package facilities

import "github.com/meso-org/meso/repository"

type Service interface {
	RegisterNewFacility(facilityName, email string) (repository.FacilityID, error)
	FindFacilityByID(repository.FacilityID) (*repository.Facility, error)
	FindAllFacilities() ([]*repository.Facility, error)
}
