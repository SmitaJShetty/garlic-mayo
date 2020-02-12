package service

//GridLocation construct for grid location
type GridLocation struct {
	X int
	Y int
}

//NewGridLocation creates a new grid location
func NewGridLocation(x int, y int) *GridLocation {
	return &GridLocation{
		X: x,
		Y: y,
	}
}
