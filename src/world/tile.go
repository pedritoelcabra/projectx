package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/container"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"github.com/pedritoelcabra/projectx/src/world/utils"
)

// Tile is the data contained in any given coordinate in the grid
type Tile struct {
	coordinates       tiling.Coord
	Data              *container.Container
	BuildingId        EntityKey
	building          *Building
	vegetation        *Vegetation
	neighbouringHexes [6]tiling.Coord
	borders           [6]bool
	hasAnyBorders     bool
	borderSprite      *ebiten.Image
}

func NewTile() *Tile {
	aTile := &Tile{}
	aTile.Data = container.NewContainer()
	return aTile
}

func (t *Tile) GetCoord() tiling.Coord {
	return t.coordinates
}

func (t *Tile) SetBuilding(building *Building) {
	t.building = building
	t.BuildingId = building.GetId()
}

func (t *Tile) GetBuilding() *Building {
	return t.building
}

func (t *Tile) SetVegetation(vegetation *Vegetation) {
	t.vegetation = vegetation
}

func (t *Tile) GetVegetation() *Vegetation {
	return t.vegetation
}

func (t *Tile) Coord() tiling.Coord {
	return t.coordinates
}

func (t *Tile) X() int {
	return t.coordinates.X()
}

func (t *Tile) Y() int {
	return t.coordinates.Y()
}

func (t *Tile) Get(key int) int {
	return t.Data.Get(key)
}

func (t *Tile) GetF(key int) float64 {
	return t.Data.GetF(key)
}

func (t *Tile) Set(key, value int) {
	t.Data.Set(key, value)
}

func (t *Tile) SetF(key int, value float64) {
	t.Data.SetF(key, value)
}

func (t *Tile) IsImpassable() bool {
	return t.GetF(MovementCost) > 100
}

func (t *Tile) IsLand() bool {
	return t.Get(TerrainBase) != utils.BasicWater && t.Get(TerrainBase) != utils.BasicDeepWater
}

func (t *Tile) CalculateMovementCost() {
	movementCost := utils.MovementCost(t.Get(TerrainBase))
	if t.vegetation != nil {
		movementCost += t.vegetation.GetMovementCost()
	}
	t.SetF(MovementCost, movementCost)
}

func (t *Tile) Recalculate() {
	t.CalculateMovementCost()
	t.borders = [6]bool{false}
	t.hasAnyBorders = false
	if t.Get(SectorId) >= 0 {
		neighbours := tiling.NeighbouringHexes(t.coordinates)
		for dir, neighbourCoord := range neighbours {
			neighbourTile := theWorld.Grid.Tile(neighbourCoord)
			if neighbourTile.Get(SectorId) != t.Get(SectorId) {
				t.borders[dir] = true
				t.hasAnyBorders = true
			}
		}
		if t.hasAnyBorders {
			imageWidth := 72  // tiling.TileHorizontalSeparation
			imageHeight := 72 // tiling.TileHeight
			t.borderSprite, _ = ebiten.NewImage(int(imageWidth), int(imageHeight), ebiten.FilterDefault)
			for dir, hasBorder := range t.borders {
				if hasBorder {
					op := &ebiten.DrawImageOptions{}
					gfx.DrawHexTerrainToImage(0, 0, utils.DirectionToBorder(dir), t.borderSprite, op)
				}
			}
		}
	}
}

func DrawSectorBorders(t *Tile) {
	if !t.hasAnyBorders {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM = ebiten.TranslateGeo(t.GetF(RenderX), t.GetF(RenderY))
	theWorld.GetScreen().DrawImage(t.borderSprite, op)
}

func DrawTerrain(t *Tile) {
	opts := &ebiten.DrawImageOptions{}
	gfx.DrawHexTerrain(t.GetF(RenderX), t.GetF(RenderY), t.Get(TerrainBase), theWorld.GetScreen(), opts)
}

func DrawVegetation(t *Tile) {
	if t.vegetation == nil {
		return
	}
	t.vegetation.Sprite.DrawSprite(theWorld.GetScreen(), t.GetF(RenderDoubleX), t.GetF(RenderDoubleY))
}

func (t *Tile) Neighbours() [6]*Tile {
	var neighbours = [6]*Tile{}
	for key, coord := range tiling.NeighbouringHexes(t.coordinates) {
		neighbours[key] = theWorld.Grid.Tile(coord)
	}
	return neighbours
}