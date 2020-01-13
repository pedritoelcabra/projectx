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
