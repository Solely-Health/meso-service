package repository

import (
	"time"

	"github.com/beevik/guid"
)

type PositionID string

type Position struct {
	PositionID PositionID
	FacilityID FacilityID
	// TODO: change to an enum Ie. Respiratory Therapist,
	Title       string
	Description string
	Schedule    []Schedule
}

type Schedule struct {
	Name  string
	Start time.Time
	End   time.Time
}

type PositionRepository interface {
	Store(position *Position) error
}

// GeneratePositionID - return a new PositionID string
func GeneratePositionID() PositionID {
	return PositionID(guid.NewString())
}
