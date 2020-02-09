package main

//PoorCreatures construct
type PoorCreature struct {
	location *GridLocation
	isZombie bool
}

//NewPoorCreature creates a new poor creature
func NewPoorCreature(x int, y int) *PoorCreature{
	return &PoorCreature{
		location: NewGridLocation(x, y),
	}
}

//GetZombiefied - what happens when a zombie eats a poor creature
func(pc *PoorCreature) GetZombiefied() {
	pc.isZombie=true
}

//GoToQueue go in queue
func(pc *PoorCreature) GoToQueue(queue *SleepingZombieQueue){
	queue.AddToQueue(pc)
}