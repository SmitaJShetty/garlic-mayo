package service

//PoorCreature construct
type PoorCreature struct {
	location *GridLocation
	isZombie bool
}

//NewPoorCreature creates a new poor creature
func NewPoorCreature(x int, y int) *PoorCreature {
	return &PoorCreature{
		location: NewGridLocation(x, y),
	}
}

//GetZombiefied - what happens when a zombie eats a poor creature
func (pc *PoorCreature) GetZombiefied(queue *SleepingZombieQueue, g *Grid) error {
	pc.isZombie = true
	err := queue.Enqueue(NewZombie(pc.location.X, pc.location.Y, g))
	return err
}
