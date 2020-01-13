package world

type FactionKey int

type Faction struct {
	Id FactionKey
}

func NewFaction() *Faction {
	aFaction := &Faction{}
	aFaction.Id = theWorld.AddFaction(aFaction)
	return aFaction
}
