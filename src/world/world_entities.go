package world

type Entities struct {
	Sectors   SectorMap
	Factions  FactionMap
	Buildings BuildingMap
	Units     UnitMap
}

func NewEntities() *Entities {
	aEntities := &Entities{}
	aEntities.Sectors = make(SectorMap)
	aEntities.Factions = make(FactionMap)
	aEntities.Buildings = make(BuildingMap)
	aEntities.Units = make(UnitMap)
	return aEntities
}

func (w *World) AddUnit(unit *Unit) UnitKey {
	key := UnitKey(len(w.Entities.Units))
	w.Entities.Units[key] = unit
	return key
}

func (w *World) GetUnit(key UnitKey) *Unit {
	if key < 0 {
		return nil
	}
	return w.Entities.Units[key]
}

func (w *World) AddBuilding(building *Building) BuildingKey {
	key := BuildingKey(len(w.Entities.Buildings))
	w.Entities.Buildings[key] = building
	return key
}

func (w *World) GetBuilding(key BuildingKey) *Building {
	if key < 0 {
		return nil
	}
	return w.Entities.Buildings[key]
}

func (w *World) AddSector(sector *Sector) SectorKey {
	key := SectorKey(len(w.Entities.Sectors))
	w.Entities.Sectors[key] = sector
	return key
}

func (w *World) GetSector(key SectorKey) *Sector {
	if key < 0 {
		return nil
	}
	return w.Entities.Sectors[key]
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

func (w *World) UnitsCollidingWith(x, y float64) UnitMap {
	collidingUnits := make(UnitMap)
	for id, unit := range w.Entities.Units {
		if unit.CollidesWith(x, y) {
			collidingUnits[id] = unit
		}
	}
	return collidingUnits
}
