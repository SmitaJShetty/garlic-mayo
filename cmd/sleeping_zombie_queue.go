package main

//SleepingZombieQueue construct zombieS
type SleepingZombieQueue struct {
	ZombieQueue []Zombie
}

//NewSleepingZombieQueue returns a new sleeping zombie queue
func NewSleepingZombieQueue() *SleepingZombieQueue {
	return &SleepingZombieQueue{
		ZombieQueue: []Zombie,
	}
}

func(sz *SleepingZombieQueue) AddToQueue(zombie *Zombie)error{

}

func (sz *SleepingZombieQueue) TakeFromQueue() (*Zombie, error){
	
}