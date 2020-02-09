package main

//Zombie construct for zombie
type Zombie struct{
	position *GridLocation
	bounds int
}

func NewZombie(x, y int, b int) *Zombie{
	return &Zombie{
		position: NewGridLocation(x,y),
		bounds: b,
	}
}

func NewZombieFromPoorCreature(pc *PoorCreature) *Zombie{
	return &Zombie{
		position: NewGridLocation(pc.bounds),
	}
}

//TraverseRoute traverses route and converts poor creatures to zombies
func(z *Zombie) TraverseRouteAndStartHavoc(mItinary string) error{
	if mItinary=="" {
		return fmt.Errorf("itinary was empty")
	}

	//start going through motions in string	
	for _,d:= range mItinary {
		switch d {
		case 'U':
			z.goUp()
			break
		case 'D':
			z.goDown() 
			break		
		case 'L':
			z.goLeft() 
			break
		case 'R':
			z.goRight()
			break
		}

		if 

	}

	//initiate motions for creatures on grid

	//start motions for items in queue
}

func(z *Zombie) goUp() error{
	z.position.Y= z.Position.Y-1
	if z.Position.Y < 0 {
		z.Position.Y= z.bounds-1
	}
}

func (z *Zombie) goDown()error{
	z.position.Y=z.position.Y+1
	if z.position.Y >= z.bounds {
		z.position.Y=0
	}
}

func(z *Zombie) goLeft()error{
	z.position.X=z.position.X-1
	if z.position.X > 0 {
		z.position.X=z.bounds -1
	}
}

func (z *Zombie) goRight() error{
	z.position.X=z.position.X+1
	if z.position.X >= z.bounds {
		z.position.Z=0
	}
}



