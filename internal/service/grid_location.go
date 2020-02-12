package service

//GridLocation construct for grid location
type GridLocation struct {
	X int `json:"x"`
	Y int  `json:"y"`
}

//NewGridLocation creates a new grid location
func NewGridLocation(x int, y int) *GridLocation {
	return &GridLocation{
		X: x,
		Y: y,
	}
}
