package service

import "log"

//Apocalypse contract for zombie service
type Apocalypse interface{
	WakeupZombieWalker(req *Request) (int, error)
}

//ZombieService service construct for zombies
type ZombieService struct {
}

//WakeupZombieWalker wake up the zombie service
func (zs *ZombieService) WakeupZombieWalker(req *Request) (int, error) {
	grid := NewGrid(Limit(req.GridDimension.X), Limit(req.GridDimension.Y), req.PoorCreatures)
	zombie := NewZombie(req.ZombieLocation.X, req.ZombieLocation.Y, grid)
	score,startErr := grid.Start(req.Itinary, zombie)
	if startErr != nil {
		log.Println("error occurred in start,", startErr)
		return 0, startErr
	}

	return score, nil
}

//NewZombieService returns a new zombieservice
func NewZombieService() Apocalypse {
	return &ZombieService{}
}
