package main

//Limit bounds
type Limit int

//Grid constrcut for grid
type Grid struct {
	XLimit Limit
	YLimit Limit
}

//NewGrid returns a new grid and sets its boundaries
func NewGrid(xLimit Limit, yLimit Limit) *Grid {
	return &Grid{
		XLimit: xLimit,
		YLimit: yLimit,
	}
}
