package core

import "github.com/pedritoelcabra/projectx/src/core/defs"

type PlacementManager struct {
	selectedBuilding *defs.BuildingDef
}

func NewPlacementManager() *PlacementManager {
	aManager := &PlacementManager{}
	return aManager
}

func (p *PlacementManager) SetBuilding(building *defs.BuildingDef) {
	p.selectedBuilding = building
}

func (p *PlacementManager) UnSetBuilding() {
	p.selectedBuilding = nil
}

func (p *PlacementManager) GetBuilding() *defs.BuildingDef {
	return p.selectedBuilding
}

func (p *PlacementManager) HasBuilding() bool {
	return p.selectedBuilding != nil
}
