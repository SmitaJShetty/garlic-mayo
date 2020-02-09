package main

func main() {
	n:=10
	morbidItinary:= "DULRDULRDULR"
	creatureLocationList:= getPoorCreatureListFromLocation()
	creatureLocationList=[]PoorCreature{
		NewPoorCreature(1,2),
		NewPoorCreature(3,3),
		NewPoorCreature(4,4),
		NewPoorCreature(2,2),
	}

	zombie:= NewZombie()
    zombie.TraverseRouteAndStartHavoc(morbidItinary)
}

//given a set of string tupes, convert to poor creature location on grid
func getPoorCreatureListFromLocation() ([]*PoorCreature,error){
	//todo
}