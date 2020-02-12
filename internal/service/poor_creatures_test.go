package service

import "testing"
import "github.com/stretchr/testify/assert"

func Test_GetZombiefied(t *testing.T) {
	sZQ := NewSleepingZombieQueue()
	x := 10
	y := 10
	pc := NewPoorCreature(x, y)
	grid := NewGrid(Limit(x), Limit(y), []*PoorCreature{pc})
	assert.NotNil(t, pc)
	err := pc.GetZombiefied(sZQ, grid)
	assert.Nil(t, err)
}
