package world

type FactionKey int
type FactionMap map[FactionKey]*Faction

const (
	RelationHostile  int = -30
	RelationFriendly int = 30
)

type Faction struct {
	Id              FactionKey
	Name            string
	Relations       map[FactionKey]int
	DefaultRelation int
}

func NewFaction(name string) *Faction {
	aFaction := &Faction{}
	aFaction.Name = name
	aFaction.Id = theWorld.AddFaction(aFaction)
	aFaction.Relations = make(map[FactionKey]int)
	aFaction.DefaultRelation = 0
	return aFaction
}

func (f *Faction) GetId() FactionKey {
	return f.Id
}

func (f *Faction) GetName() string {
	return f.Name
}

func (f *Faction) GetRelation(f2 *Faction) int {
	if rel, ok := f.Relations[f2.Id]; ok {
		return rel
	}
	f.Relations[f2.Id] = f.DefaultRelation
	return f.DefaultRelation
}

func RelationIsHostile(level int) bool {
	return level <= RelationHostile
}

func RelationIsFriendly(level int) bool {
	return level >= RelationFriendly
}
