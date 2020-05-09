package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/core/defs"
	container2 "github.com/pedritoelcabra/projectx/src/core/world/container"
	tiling2 "github.com/pedritoelcabra/projectx/src/core/world/tiling"
	utils2 "github.com/pedritoelcabra/projectx/src/core/world/utils"
	"github.com/pedritoelcabra/projectx/src/gfx"
)

// Tile is the data contained in any given coordinate in the grid
type Tile struct {
	coordinates       tiling2.Coord
	Data              *container2.Container
	Building          BuildingPointer
	neighbouringHexes [6]tiling2.Coord
	borders           [6]bool
	hasAnyBorders     bool
	borderSprite      *ebiten.Image
}

func NewTile() *Tile {
	aTile := &Tile{}
	aTile.Building = MakeEmptyBuildingPointer()
	aTile.Data = container2.NewContainer()
	return aTile
}

func (t *Tile) GetCoord() tiling2.Coord {
	return t.coordinates
}

func (t *Tile) GetRenderPos() (x, y float64) {
	return t.GetF(RenderX), t.GetF(RenderY)
}

func (t *Tile) GetCenterPos() (x, y float64) {
	return t.GetF(CenterX), t.GetF(CenterY)
}

func (t *Tile) SetBuilding(building *Building) {
	t.Building = building.GetPointer()
}

func (t *Tile) GetBuilding() *Building {
	return t.Building.Get()
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
	return t.Get(TerrainBase) != utils2.BasicWater && t.Get(TerrainBase) != utils2.BasicDeepWater
}

func (t *Tile) CalculateMovementCost() {
	movementCost := utils2.MovementCost(t.Get(TerrainBase))
	resource := t.Get(Resource)
	if resource != 0 {
		movementCost += defs.ResourceById(resource).MovementCost
	}
	t.SetF(MovementCost, movementCost)
}

func (t *Tile) Recalculate() {
	t.CalculateMovementCost()
	t.borders = [6]bool{false}
	t.hasAnyBorders = false
	if t.HasSector() {
		neighbours := tiling2.NeighbouringHexes(t.coordinates)
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
					gfx.DrawHexTerrainToImage(0, 0, utils2.DirectionToBorder(dir), t.borderSprite, op)
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

func DrawResource(t *Tile) {
	if t.Get(Resource) != 0 {
		defs.DrawResource(t.Get(Resource), theWorld.GetScreen(), t.GetF(RenderDoubleX), t.GetF(RenderDoubleY))
	}
}

func (t *Tile) Neighbours() [6]*Tile {
	var neighbours = [6]*Tile{}
	for key, coord := range tiling2.NeighbouringHexes(t.coordinates) {
		neighbours[key] = theWorld.Grid.Tile(coord)
	}
	return neighbours
}

func (t *Tile) HasSector() bool {
	return t.Get(SectorId) >= 0
}

func (t *Tile) GetSector() *Sector {
	return theWorld.GetSector(SectorKey(t.Get(SectorId)))
}

func (t *Tile) GetFaction() *Faction {
	return theWorld.GetFaction(FactionKey(t.Get(FactionId)))
}

func (t *Tile) OwnedByPlayer() bool {
	sector := t.GetSector()
	if sector == nil {
		return false
	}
	return sector.GetFaction().Id == theWorld.PlayerUnit.GetFaction().Id
}

func (t *Tile) GetSectorId() SectorKey {
	sector := t.GetSector()
	if sector == nil {
		return -1
	}
	return sector.Id
}

func (t *Tile) SpawnUnit(name string) *Unit {
	return NewUnit(name, tiling2.NewCoordF(t.GetRenderPos()))
}

func (t *Tile) HasResource(name string) bool {
	resource := t.Get(Resource)
	if resource == 0 {
		return false
	}
	return name == defs.ResourceById(resource).Resource
}

func (t *Tile) GetResource() int {
	return t.Get(Resource)
}

func (t *Tile) GetResourceAmount() int {
	return t.Get(ResourceAmount)
}
