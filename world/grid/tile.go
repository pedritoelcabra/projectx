package grid

import "github.com/pedritoelcabra/projectx/world/coord"

// Tile is the data contained in any given coordinate in the grid
type Tile struct {
	coordinates coord.Coord
	values      map[int]int
	valuesF     map[int]float64
}

func (t Tile) X() int {
	return t.coordinates.X()
}

func (t Tile) Y() int {
	return t.coordinates.Y()
}

func (t Tile) Get(key int) int {
	return t.values[key]
}

func (t Tile) GetF(key int) float64 {
	return t.valuesF[key]
}

func (t Tile) Set(key, value int) {
	t.values[key] = value
}

func (t Tile) SetF(key int, value float64) {
	t.valuesF[key] = value
}

const (
	Height int = iota
	Population
	TerrainBase
	TerrainOverlay
	RenderX
	RenderY
	CenterX
	CenterY
	RenderW
	RenderH
)
