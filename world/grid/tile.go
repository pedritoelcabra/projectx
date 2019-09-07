package grid

// Tile is the data contained in any given coordinate in the grid
type Tile struct {
	coordinates Coord
	values      map[int]int
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

func (t Tile) Set(key, value int) {
	t.values[key] = value
}

const (
	Height int = iota
	Population
	TerrainBase
	TerrainOverlay
)
