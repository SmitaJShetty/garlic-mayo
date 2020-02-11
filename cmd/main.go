package main

import "log"

func main() {
	n := 10
	morbidItinary := "DULRDULRDULR"
	//creatureLocationList:= getPoorCreatureListFromLocation()
	creatureLocationList := []*PoorCreature{
		NewPoorCreature(1, 2),
		NewPoorCreature(3, 3),
		NewPoorCreature(4, 4),
		NewPoorCreature(2, 2),
	}

	grid := NewGrid(Limit(n), Limit(n), creatureLocationList)
	zombie := NewZombie(2, 1, grid)
	startErr := grid.Start(morbidItinary, zombie)
	if startErr != nil {
		log.Println("error occurred in start,", startErr)
	}
}

//given a set of string tupes, convert to poor creature location on grid
func getPoorCreatureListFromLocation() ([]*PoorCreature, error) {
	//todo
	return nil, nil
}
