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

//AddToQueue adds zombie to queue
func(sz *SleepingZombieQueue) AddToQueue(zombie *Zombie)error{
	
}

//TakeFromQueue takes out of sleeping zombie queue
func (sz *SleepingZombieQueue) TakeFromQueue() (*Zombie, error){

}