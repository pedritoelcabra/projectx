package world

import (
	"github.com/pedritoelcabra/projectx/world/container"
	"github.com/pedritoelcabra/projectx/world/tiling"
	"github.com/pedritoelcabra/projectx/world/utils"
)

// Tile is the data contained in any given coordinate in the grid
type Tile struct {
	coordinates       tiling.Coord
	Data              *container.Container
	BuildingId        EntityKey
	building          *Building
	neighbouringHexes [6]tiling.Coord
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

func (t *Tile) Recalculate() {
	t.SetF(MovementCost, utils.MovementCost(t.Get(TerrainBase)))
}
