package world

type WorldEntities struct {
	Entities  EntityMap
	Sectors   SectorMap
	Factions  FactionMap
	Buildings BuildingMap
}

func NewWorldEntities() *WorldEntities {
	aEntities := &WorldEntities{}
	aEntities.Entities = make(EntityMap)
	aEntities.Sectors = make(SectorMap)
	aEntities.Factions = make(FactionMap)
	aEntities.Buildings = make(BuildingMap)
	return aEntities
}
