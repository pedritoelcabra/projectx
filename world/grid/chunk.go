package grid

type chunk struct {
	tiles    []*tile
	location Coord
}

func NewChunk(location Coord) *chunk {
	arrayTiles := make([]*tile, ChunkSize*ChunkSize)
	aChunk := &chunk{arrayTiles, location}
	for x := 0; x < ChunkSize; x++ {
		for y := 0; y < ChunkSize; y++ {
			tileX := (location.X() * ChunkSize) + x
			tileY := (location.Y() * ChunkSize) + y
			tileLocation := NewCoord(tileX, tileY)
			tileIndex := aChunk.tileIndex(tileX, tileY)
			aTile := &tile{tileLocation, make(map[int]int)}
			aChunk.tiles[tileIndex] = aTile
		}
	}
	return aChunk
}

func (ch *chunk) RunOnAllTiles(f func(t *tile)) {
	for _, c := range ch.tiles {
		f(c)
	}
}

func (ch *chunk) Tile(tileCoord Coord) *tile {
	return ch.tiles[ch.tileIndex(tileCoord.X(), tileCoord.Y())]

}

func (ch *chunk) initTile(c Coord) *tile {
	return &tile{c, make(map[int]int)}
}

func (ch *chunk) tileIndex(x, y int) int {
	x -= ch.location.X() * ChunkSize
	y -= ch.location.Y() * ChunkSize
	return (x * ChunkSize) + y
}
