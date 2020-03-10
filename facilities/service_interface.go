package facilities

import "github.com/meso-org/meso/repository"

type Service interface {
	RegisterNewFacility() (repository.FacilityID, error)
}
