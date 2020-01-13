package world

type WorldEntities struct {
	Sectors   SectorMap
	Factions  FactionMap
	Buildings BuildingMap
	Units     UnitMap
}

func NewWorldEntities() *WorldEntities {
	aEntities := &WorldEntities{}
	aEntities.Sectors = make(SectorMap)
	aEntities.Factions = make(FactionMap)
	aEntities.Buildings = make(BuildingMap)
	aEntities.Units = make(UnitMap)
	return aEntities
}
