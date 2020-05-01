package repository

import (
	"fmt"
	"time"

	"github.com/meso-org/meso/config"

	"github.com/beevik/guid"
)

//TODO: make this live in global somewhere
type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(config.Dateformat))
	return []byte(stamp), nil
}

type PositionID string

type Position struct {
	PositionID PositionID
	FacilityID FacilityID
	// TODO: change to an enum Ie. Respiratory Therapist,
	Title         string
	Description   string
	StartDateTime JSONTime
	EndDateTime   JSONTime
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
