package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TraverseZombiePath_Success(t *testing.T) {
	sZQ := NewSleepingZombieQueue()
	x := 10
	y := 10
	pc := NewPoorCreature(1, 1)
	grid := NewGrid(Limit(x), Limit(y), []*PoorCreature{pc})

	zombie := NewZombie(0, 1, grid)
	itinary := "RRRR"
	err := zombie.TraverseZombiePath(itinary, sZQ)
	assert.Nil(t, err)
	assert.Equal(t, 4, zombie.position.X)
	assert.Equal(t, 1, zombie.position.Y)
	assert.Equal(t, 1, len(sZQ.ZombieQueue))
}

func Test_TraverseZombiePath_Error(t *testing.T) {
	sZQ := NewSleepingZombieQueue()
	x := 10
	y := 10
	pc := NewPoorCreature(x, y)
	grid := NewGrid(Limit(x), Limit(y), []*PoorCreature{pc})

	zombie := NewZombie(1, 2, grid)
	itinary := ""
	err := zombie.TraverseZombiePath(itinary, sZQ)
	assert.Nil(t, err)
}

func Test_NewZombieFromPoorCreature_Success(t *testing.T) {
	x := 10
	y := 10
	pc := NewPoorCreature(x, y)
	grid := NewGrid(Limit(x), Limit(y), []*PoorCreature{pc})

	zombie := NewZombieFromPoorCreature(pc, grid)
	assert.NotNil(t, zombie)
	assert.NotNil(t, zombie.grid)
	assert.Equal(t, Limit(10), zombie.grid.XLimit)
}

func Test_goUp_Success(t *testing.T) {
	x := 10
	y := 11
	grid := NewGrid(Limit(x), Limit(y), []*PoorCreature{})

	zombie := NewZombie(10, 11, grid)
	zombie.goUp()
	assert.Equal(t, 10, zombie.position.X)
	assert.Equal(t, 10, zombie.position.Y)
}

func Test_goUp_Success_Wrapup(t *testing.T) {
	x := 10
	y := 10
	grid := NewGrid(Limit(x), Limit(y), []*PoorCreature{})

	zombie := NewZombie(8, 0, grid)
	zombie.goUp()
	assert.Equal(t, 8, zombie.position.X)
	assert.Equal(t, 9, zombie.position.Y)
}

func Test_goDown_Success(t *testing.T) {
	x := 10
	y := 11
	grid := NewGrid(Limit(x), Limit(y), []*PoorCreature{})
	zombie := NewZombie(10, 10, grid)
	zombie.goDown()
	assert.Equal(t, 10, zombie.position.X)
	assert.Equal(t, 0, zombie.position.Y)
}

func Test_goLeft_Success(t *testing.T) {
	x := 10
	y := 11
	grid := NewGrid(Limit(x), Limit(y), []*PoorCreature{})
	zombie := NewZombie(10, 11, grid)
	zombie.goLeft()
	assert.Equal(t, 9, zombie.position.X)
	assert.Equal(t, 11, zombie.position.Y)
}

func Test_goRight_Success(t *testing.T) {
	x := 10
	y := 11
	grid := NewGrid(Limit(x), Limit(y), []*PoorCreature{})
	zombie := NewZombie(x, y, grid)
	zombie.goRight()
	assert.Equal(t, 0, zombie.position.X)
	assert.Equal(t, 11, zombie.position.Y)
}

func Test_goRight_Success_Wrapup(t *testing.T) {
	x := 10
	y := 10
	grid := NewGrid(Limit(x), Limit(y), []*PoorCreature{})
	zombie := NewZombie(x, y, grid)
	zombie.goRight()
	assert.Equal(t, 0, zombie.position.X)
	assert.Equal(t, 10, zombie.position.Y)
}
