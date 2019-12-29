package world

import (
	"github.com/pedritoelcabra/projectx/world/container"
	"github.com/pedritoelcabra/projectx/world/tiling"
	"github.com/pedritoelcabra/projectx/world/utils"
)

// Tile is the data contained in any given coordinate in the grid
type Tile struct {
	coordinates tiling.Coord
	Data        *container.Container
}

func NewTile() *Tile {
	aTile := &Tile{}
	aTile.Data = container.NewContainer()
	return aTile
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
	return t.MovementCost() == 0
}

func (t *Tile) MovementCost() float64 {
	return utils.MovementCost(t.Get(TerrainBase))
}
