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
	startRelation := f.DefaultRelation
	if f.Id == f2.Id {
		startRelation = 100
	}
	f.Relations[f2.Id] = startRelation
	return startRelation
}

func (f *Faction) IsHostileTowards(f2 *Faction) bool {
	return RelationIsHostile(f.GetRelation(f2))
}

func (f *Faction) IsFriendlyTowards(f2 *Faction) bool {
	return RelationIsHostile(f.GetRelation(f2))
}

func RelationIsHostile(level int) bool {
	return level <= RelationHostile
}

func RelationIsFriendly(level int) bool {
	return level >= RelationFriendly
}
