package facilities

import (
	"fmt"

	"github.com/meso-org/meso/repository"
)

type service struct {
	facility repository.FacilityRepository
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

func (s *service) FindFacilityByID(id repository.FacilityID) (*repository.Facility, error) {

	if id == "" {
		return nil, err
	}

	facility, err := s.facilities.Find(id)
	if err != nil {
		return nil, err
	}

	return facility, nil
}
