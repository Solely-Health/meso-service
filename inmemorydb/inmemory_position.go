package inmemorydb

import (
	"fmt"
	"sync"

	repository "github.com/meso-org/meso/repository"
)

type positionRepository struct {
	mtx       sync.RWMutex
	positions map[repository.PositionID]*repository.Position
}

func (r *positionRepository) Store(p *repository.Position) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.positions[p.PositionID] = p
	if r.positions[p.PositionID] != p {
		return fmt.Errorf("seems as if in memory store errored out")
	}
	return nil
}

func (r *positionRepository) Find(x interface{}) (*repository.Position, error) {
	switch x.(type) {
	case repository.PositionID:
		id := repository.PositionID(fmt.Sprintf("%v", x))
		r.mtx.Lock()
		defer r.mtx.Unlock()
		position := r.positions[id]
		if position == nil {
			return position, fmt.Errorf("Could not find position by id: %v", id)
		}
		return position, nil
	default:
		return nil, fmt.Errorf("Cannot find position, bad parameter type")
	}
}

func (r *positionRepository) FindByFacilityID(facilityID repository.FacilityID) ([]*repository.Position, error) {
	var positions []*repository.Position
	for _, position := range r.positions {
		if position.FacilityID == facilityID {
			positions = append(positions, position)
		}
	}
	return positions, nil
}

func (r *positionRepository) FindAll() ([]*repository.Position, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	positions := []*repository.Position{}
	for _, position := range r.positions {
		positions = append(positions, position)
	}

	return positions, nil
}

// NewPositionRepository returns a new instance of a in-memory cargo repository.
func NewPositionRepository() repository.PositionRepository {
	return &positionRepository{
		positions: make(map[repository.PositionID]*repository.Position),
	}
}
