package service

import (
	"fmt"
	"log"
	"strconv"
)

//Zombie construct for zombie
type Zombie struct {
	position *GridLocation
	grid     *Grid
}

//NewZombie construct for new zombie
func NewZombie(x, y int, g *Grid) *Zombie {
	return &Zombie{
		position: NewGridLocation(x, y),
		grid:     g,
	}
}

//NewZombieFromPoorCreature returns a new zombie from poor creature
func NewZombieFromPoorCreature(pc *PoorCreature, g *Grid) *Zombie {
	return &Zombie{
		position: NewGridLocation(pc.location.X, pc.location.Y),
		grid:     g,
	}
}

//TraverseZombiePath walk in the zombie footsteps
func (z *Zombie) TraverseZombiePath(mItinary string, zombieQueue *SleepingZombieQueue) error {
	latestItinary := make(map[string]int)
	//start going through motions in string
	for _, d := range mItinary {
		switch d {
		case 'U':
			z.goUp()
			break
		case 'D':
			z.goDown()
			break
		case 'L':
			z.goLeft()
			break
		case 'R':
			z.goRight()
			break
		}

		latestItinary[markerStr(z.position.X, z.position.Y)] = 1
	}

	if len(z.grid.PoorCreatures) == 0 {
		return fmt.Errorf("no poor creatures")
	}

	log.Println("pcs:", z.grid.PoorCreatures[1])
	for _, pc := range z.grid.PoorCreatures {
		if !pc.isZombie {
			_, ok := latestItinary[markerStr(pc.location.X, pc.location.Y)]
			if ok {
				//poor creature lived here
				err := pc.GetZombiefied(zombieQueue, z.grid)
				if err != nil {
					log.Println("error occurred when pc got zombified")
				}
			}
		}
	}
	return nil
}

func markerStr(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func (z *Zombie) goUp() error {
	z.position.Y = z.position.Y - 1
	if z.position.Y < 0 {
		z.position.Y = int(z.grid.YLimit) - 1
	}

	return nil
}

func (z *Zombie) goDown() error {
	z.position.Y = z.position.Y + 1
	if z.position.Y >= int(z.grid.YLimit) {
		z.position.Y = 0
	}

	return nil
}

func (z *Zombie) goLeft() error {
	z.position.X = z.position.X - 1
	if z.position.X > 0 {
		z.position.X = int(z.grid.XLimit) - 1
	}

	return nil
}

func (z *Zombie) goRight() error {
	z.position.X = z.position.X + 1
	if z.position.X >= int(z.grid.XLimit) {
		z.position.X = 0
	}
	return nil
}
