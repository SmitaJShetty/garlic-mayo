package main

//SleepingZombieQueue construct zombieS
type SleepingZombieQueue struct {
	ZombieQueue []*Zombie
}

//NewSleepingZombieQueue returns a new sleeping zombie queue
func NewSleepingZombieQueue() *SleepingZombieQueue {
	return &SleepingZombieQueue{
		
	}
}

//Enqueue adds zombie to queue
func(sz *SleepingZombieQueue) Enqueue(zombie *Zombie)error{
		return nil
}

//Dequeue takes out of sleeping zombie queue
func (sz *SleepingZombieQueue) Dequeue() (*Zombie, error){
	return nil, nil
}