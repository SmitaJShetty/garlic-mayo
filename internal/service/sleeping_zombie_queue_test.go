package service

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Enqueue(t *testing.T) {
	sZQ := NewSleepingZombieQueue()
	x := 10
	y := 10
	grid := NewGrid(Limit(x), Limit(y), []*PoorCreature{})
	zombie := NewZombie(1, 2, grid)
	err := sZQ.Enqueue(zombie)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(sZQ.ZombieQueue))
}

func Test_Dequeue(t *testing.T) {
	sZQ := NewSleepingZombieQueue()
	x := 10
	y := 10
	grid := NewGrid(Limit(x), Limit(y), []*PoorCreature{})
	zombie := NewZombie(1, 2, grid)
	sZQ.ZombieQueue = append(sZQ.ZombieQueue, zombie)
	assert.NotNil(t, 1, len(sZQ.ZombieQueue))

	actualZombie, err := sZQ.Dequeue()
	assert.Nil(t, err)
	assert.NotNil(t, actualZombie)
}
