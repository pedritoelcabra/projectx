package grid

// Tile is the data contained in any given coordinate in the grid
type tile struct {
	coordinates Coord
	values      map[int]int
}

func (t tile) X() int {
	return t.coordinates.X()
}

func (t tile) Y() int {
	return t.coordinates.Y()
}

func (t tile) Get(key int) int {
	return t.values[key]
}

func (t tile) Set(key, value int) {
	t.values[key] = value
}

const (
	Height int = iota
	Population
	TerrainBase
	TerrainOverlay
)
