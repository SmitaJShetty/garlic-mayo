package service

//SleepingZombieQueue construct zombieS
type SleepingZombieQueue struct {
	ZombieQueue []*Zombie
}

//NewSleepingZombieQueue returns a new sleeping zombie queue
func NewSleepingZombieQueue() *SleepingZombieQueue {
	return &SleepingZombieQueue{
		ZombieQueue: []*Zombie{},
	}
}

//Enqueue adds zombie to queue
func (sz *SleepingZombieQueue) Enqueue(zombie *Zombie) error {
	sz.ZombieQueue = append(sz.ZombieQueue, zombie)
	return nil
}

//Dequeue takes out of sleeping zombie queue
func (sz *SleepingZombieQueue) Dequeue() (*Zombie, error) {
	zombie := sz.ZombieQueue[0]
	sz.ZombieQueue[0] = nil
	sz.ZombieQueue = sz.ZombieQueue[1:]
	return zombie, nil
}
