package main

import "fmt"

//Limit bounds
type Limit int

//Grid construct for grid
type Grid struct {
	XLimit        Limit
	YLimit        Limit
	PoorCreatures []*PoorCreature
}

//NewGrid returns a new grid and sets its boundaries
func NewGrid(xLimit Limit, yLimit Limit, pc []*PoorCreature) *Grid {
	return &Grid{
		XLimit:        xLimit,
		YLimit:        yLimit,
		PoorCreatures: pc,
	}
}

//Start starting place for all action
func (g *Grid) Start(mItinary string, initiatorZombie *Zombie) error {
	zombieQueue := NewSleepingZombieQueue()
	err := zombieQueue.Enqueue(initiatorZombie)
	if err != nil {
		return err
	}

	if mItinary == "" {
		return fmt.Errorf("itinary was empty")
	}

	for {
		//initiate motions for creatures on grid
		zombie, zErr := zombieQueue.Dequeue()
		if zErr != nil {
			return zErr
		}

		//traverse zombiepath for as long as there are zombies in the queue
		traverseErr := zombie.TraverseZombiePath(mItinary, zombieQueue)
		if traverseErr != nil {
			return traverseErr
		}
	}
	return nil

}
