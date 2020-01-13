package world

type FactionKey int
type FactionMap map[FactionKey]*Faction

type Faction struct {
	Id FactionKey
}

func NewFaction() *Faction {
	aFaction := &Faction{}
	aFaction.Id = theWorld.AddFaction(aFaction)
	return aFaction
}

func (f *Faction) GetId() FactionKey {
	return f.Id
}

func (w *World) AddFaction(sector *Faction) FactionKey {
	key := FactionKey(len(w.Entities.Factions))
	w.Entities.Factions[key] = sector
	return key
}

func (w *World) GetFaction(key FactionKey) *Faction {
	if key < 0 {
		return nil
	}
	return w.Entities.Factions[key]
}
