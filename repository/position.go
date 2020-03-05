package repository

import (
	"github.com/beevik/guid"
)

type PositionID string

type Position struct {
	PositionID PositionID
	FacilityID FacilityID
	// TODO: change to an enum Ie. Respiratory Therapist,
	Title       string
	Description string
}

type PositionRepository interface {
	Store(position *Position)
}

// GeneratePositionID - return a new PositionID string
func GeneratePositionID() PositionID {
	return PositionID(guid.NewString())
}
