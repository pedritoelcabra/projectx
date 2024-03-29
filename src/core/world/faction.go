package world

import "github.com/pedritoelcabra/projectx/src/core/defs"

type FactionKey int
type FactionMap map[FactionKey]*Faction

const (
	RelationHostile  int = -30
	RelationFriendly int = 30
)

const (
	DefaultMonsterFactionName string = "Monsters"
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
	def := defs.GetFactionDef(name)
	if def != nil {
		aFaction.DefaultRelation = def.DefaultRelation
		aFaction.Name = def.Name
	}
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
	lowestRelation := f.DefaultRelation
	if f2.DefaultRelation < lowestRelation {
		lowestRelation = f2.DefaultRelation
	}
	if f.Id == f2.Id {
		lowestRelation = 100
	}
	f.SetRelation(f2, lowestRelation)
	f2.SetRelation(f, lowestRelation)
	return lowestRelation
}

func (f *Faction) SetRelation(f2 *Faction, value int) {
	f.Relations[f2.Id] = value
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
