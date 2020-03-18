package facilities

import (
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/meso-org/meso/inmemorydb"
	"github.com/meso-org/meso/repository"
)

func TestCreateJob(t *testing.T) {
	// mock variables
	var (
		positionID = repository.GeneratePositionID()
		facilityID = repository.GenerateFacilityID()
		desc       = "This is a test desc"
		title      = "This is a Test title"
		schedules  = []repository.Schedule{
			{Name: "Schedule 1",
				Start: time.Now(),
				End:   time.Now().Add(time.Duration(6000))},
			{Name: "Schedule 2",
				Start: time.Now(),
				End:   time.Now().Add(time.Duration(6000))},
		}
	)

	shouldBePosition := repository.Position{
		PositionID:  positionID,
		FacilityID:  facilityID,
		Title:       title,
		Description: desc,
		Schedule:    schedules,
	}

	mockPositionRepo := inmemorydb.NewPositionRepository()
	mockFacilitiyRepo := inmemorydb.NewFacilityRepo()

	mockService := NewService(mockFacilitiyRepo, mockPositionRepo)

	pos, err := mockService.CreateJob(facilityID, positionID, schedules, title, desc)

	if err != nil {
		t.Fatalf("Failed test TestCreateJob: %v", err)
	}

	if diff := deep.Equal(shouldBePosition, pos); diff != nil {
		t.Fatalf("Return statement failed to be accurate: %v", diff)
	}
}
