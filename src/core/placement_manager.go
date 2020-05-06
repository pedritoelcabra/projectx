package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/world"
	"github.com/pedritoelcabra/projectx/src/gfx"
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
	tile := ProjectX.CurrentMouseTile()
	if tile == nil {
		return
	}
	x := tile.GetF(world.RenderX)
	y := tile.GetF(world.RenderY)
	op := &ebiten.DrawImageOptions{}
	red := 1.0
	if !p.BuildingCanBePlacedAtTile(tile) {
		red = 3.0
	}
	op.ColorM.Scale(red, 1.0, 1.0, 0.7)
	p.sprite.DrawSpriteWithOptions(screen, x, y, op)
}

func (p *PlacementManager) BuildingCanBePlacedAtTile(tile *world.Tile) bool {
	if tile.IsImpassable() {
		return false
	}
	if !tile.OwnedByPlayer() {
		return false
	}
	if tile.GetBuilding() != nil {
		return false
	}
	return true
}

func (p *PlacementManager) PlaceBuilding() {
	if !p.HasBuilding() {
		return
	}
	tile := ProjectX.CurrentMouseTile()
	if tile == nil {
		return
	}
	if !p.BuildingCanBePlacedAtTile(tile) {
		return
	}
	building := world.NewBuilding(p.selectedBuilding.Name, tile)
	building.StartConstruction()
	_ = building
}
