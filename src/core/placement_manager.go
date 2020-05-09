package core

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/world"
	"github.com/pedritoelcabra/projectx/src/core/world/tiling"
	"github.com/pedritoelcabra/projectx/src/core/world/utils"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/gui"
	"image"
	"image/color"
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
	p.DrawGatheringInfo(tile)
}

func (p *PlacementManager) DrawGatheringInfo(tile *world.Tile) {
	if p.selectedBuilding.Gathers == "" {
		return
	}
	radius := p.selectedBuilding.GatherRadius
	coordCenter := tile.GetCoord()
	gatherTiles := 0
	for x := tile.X() - radius; x <= tile.X()+radius; x++ {
		for y := tile.Y() - radius; y <= tile.Y()+radius; y++ {
			tileCoord := tiling.NewCoord(x, y)
			if tileCoord.Equals(coordCenter) {
				continue
			}
			distance := int(tiling.HexDistance(coordCenter, tileCoord))
			if distance > radius {
				continue
			}
			t := ProjectX.World.Grid.Tile(tileCoord)
			op := &ebiten.DrawImageOptions{}
			op.ColorM.Scale(1.0, 1.0, 1.0, 0.5)
			color := utils.GreenOverlay
			if !t.HasResource(p.selectedBuilding.Gathers) {
				color = utils.RedOverlay
			}
			sector := t.GetSector()
			if sector == nil {
				color = utils.RedOverlay
			}
			building := t.GetBuilding()
			if building != nil {
				color = utils.RedOverlay
			}
			if color == utils.GreenOverlay {
				gatherTiles++
			}
			gfx.DrawHexTerrain(t.GetF(world.RenderX), t.GetF(world.RenderY), color, ProjectX.World.GetScreen(), op)
		}
	}
	percentage := world.GetGatheringEfficiency(p.selectedBuilding, gatherTiles)
	percentageString := fmt.Sprintf("%.0f", percentage) + " %"
	xc, yc := tile.GetCenterPos()
	xci := int(xc) - 20
	yci := int(yc) - 20
	rect := image.Rect(xci, yci, xci+50, yci+50)
	DrawTextBoxOnWorldPos(percentageString, rect)
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

func DrawTextBoxOnWorldPos(text string, box image.Rectangle) {
	aBox := gui.NewTextBox()
	aBox.SetBox(box)
	aBox.SetText(text)
	aBox.SetColor(color.White)
	aBox.BuildTextBoxImage(ProjectX.Gui, box)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(box.Min.X), float64(box.Min.Y))
	ProjectX.Screen.DrawImage(aBox.GetContentBuffer(), op)
}
