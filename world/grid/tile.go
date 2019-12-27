package grid

import "github.com/pedritoelcabra/projectx/world/coord"

// Tile is the data contained in any given coordinate in the grid
type Tile struct {
	coordinates coord.Coord
	Values      map[int]int
	ValuesF     map[int]float64
}

func NewTile() *Tile {
	aTile := &Tile{}
	aTile.Values = make(map[int]int)
	aTile.ValuesF = make(map[int]float64)
	return aTile
}

func (t Tile) X() int {
	return t.coordinates.X()
}

func (t Tile) Y() int {
	return t.coordinates.Y()
}

func (t Tile) Get(key int) int {
	return t.Values[key]
}

func (t Tile) GetF(key int) float64 {
	return t.ValuesF[key]
}

func (t Tile) Set(key, value int) {
	t.Values[key] = value
}

func (t Tile) SetF(key int, value float64) {
	t.ValuesF[key] = value
}
