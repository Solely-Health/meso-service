package repository

import (
	"time"

	"github.com/beevik/guid"
)

type PositionID string
type StartDateTime time.Time
type EndDateTime time.Time

type Position struct {
	PositionID PositionID
	FacilityID FacilityID
	// TODO: change to an enum Ie. Respiratory Therapist,
	Title         string
	Description   string
	StartDateTime StartDateTime
	EndDateTime   EndDateTime
}

type PositionRepository interface {
	Store(position *Position) error
	Find(x interface{}) (*Position, error)
	FindAll() ([]*Position, error)
}

// GeneratePositionID - return a new PositionID string
func GeneratePositionID() PositionID {
	return PositionID(guid.NewString())
}
