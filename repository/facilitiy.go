package repository

import (
	"github.com/beevik/guid"
)

type FacilityID string

type FacilityRepository interface {
	Store(facility *Facility) error
	Find(x interface{}) (*Facility, error)
	FindAll() ([]*Facility, error)
}

// Facility - domain object
type Facility struct {
	FacilityID   FacilityID
	FacilityName string
	Email        string
}

// NewFacility - generate a new facility domain objec with provided fields
func NewFacility(facilityID FacilityID, facilityName, email string) *Facility {
	return &Facility{
		FacilityID:   facilityID,
		FacilityName: facilityName,
		Email:        email,
	}
}

// GenerateFacilityID -
func GenerateFacilityID() FacilityID {
	return FacilityID(guid.NewString())
}
