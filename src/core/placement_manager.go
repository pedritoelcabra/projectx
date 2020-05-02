package core

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world"
)

type PlacementManager struct {
	selectedBuilding *defs.BuildingDef
	sprite           gfx.Sprite
}

func NewPlacementManager() *PlacementManager {
	aManager := &PlacementManager{}
	return aManager
}

func (p *PlacementManager) SetBuilding(building *defs.BuildingDef) {
	p.selectedBuilding = building
	key := gfx.GetSpriteKey(building.Graphic)
	p.sprite = gfx.NewHexSprite(key)
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

func (p *PlacementManager) Draw(screen *gfx.Screen) {
	if !p.HasBuilding() {
		return
	}
	pos := ProjectX.MouseTileCoord()
	tile := ProjectX.World.Grid.Tile(pos)
	x := tile.GetF(world.RenderX)
	y := tile.GetF(world.RenderY)
	p.sprite.DrawSprite(screen, x, y)
}
