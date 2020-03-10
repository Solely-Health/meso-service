package facilities

import (
	"github.com/meso-org/meso/repository"
)

type mockFacilitiesRepository struct {
	facility *repository.Facility
}

func (mockFacility *mockFacilitiesRepository) Store(facility *repository.Facility) error {
	mockFacility.facility = facility
	return nil
}
