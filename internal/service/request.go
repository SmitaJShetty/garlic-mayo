package service

//Request model for request
type Request struct {
	Itinary        string          `json:"itinary"`
	GridDimension  GridDimension   `json:"grid_dimension"`
	PoorCreatures  []*PoorCreature `json:"poor_creatures"`
	ZombieLocation *GridLocation   `json:"zombie_location"`
}

//GridDimension construct grid dimension
type GridDimension struct {
	X int `json:"x"`
	Y int `json:"y"`
}
