package service

import "fmt"


//Request model for request
type Request struct {
	Itinary        string          `json:"itinary"`
	GridDimension  *GridDimension  `json:"grid_dimension"`
	PoorCreatures  []*PoorCreature `json:"poor_creatures"`
	ZombieLocation *GridLocation   `json:"zombie_location"`
}

//GridDimension construct grid dimension
type GridDimension struct {
	X int `json:"x"`
	Y int `json:"y"`
}

//Validate validates
func (r *Request) Validate() error {
	if r.Itinary == "" {
		return fmt.Errorf("itinary was empty")
	}

	if r.GridDimension == nil {
		return fmt.Errorf("grid dimension was empty")
	}

	if r.PoorCreatures == nil || len(r.PoorCreatures) == 0 {
		return fmt.Errorf("poor creatures was not provided")
	}

	if r.ZombieLocation == nil {
		return fmt.Errorf("zombie's location was not provided")
	}

	return nil
}
