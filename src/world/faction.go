package world

type FactionKey int
type FactionMap map[FactionKey]*Faction

type Faction struct {
	Id   FactionKey
	Name string
}

func NewFaction(name string) *Faction {
	aFaction := &Faction{}
	aFaction.Name = name
	aFaction.Id = theWorld.AddFaction(aFaction)
	return aFaction
}

func (f *Faction) GetId() FactionKey {
	return f.Id
}

func (f *Faction) GetName() string {
	return f.Name
}