package grid

type chunk struct {
	tiles               []*Tile
	location            Coord
	Generated           bool
	queuedForGeneration bool
}

func NewChunk(location Coord) *chunk {
	aChunk := &chunk{}
	aChunk.tiles = make([]*Tile, ChunkSize*ChunkSize)
	aChunk.location = location
	aChunk.Generated = false
	aChunk.queuedForGeneration = false
	for x := 0; x < ChunkSize; x++ {
		for y := 0; y < ChunkSize; y++ {
			tileX := (location.X() * ChunkSize) + x
			tileY := (location.Y() * ChunkSize) + y
			tileLocation := NewCoord(tileX, tileY)
			tileIndex := aChunk.tileIndex(tileX, tileY)
			aTile := &Tile{tileLocation, make(map[int]int)}
			aChunk.tiles[tileIndex] = aTile
		}
	}
	return aChunk
}

func (ch *chunk) IsGenerated() bool {
	return ch.Generated
}

func (ch *chunk) IsQueueForGeneration() bool {
	return ch.queuedForGeneration
}

func (ch *chunk) RunOnAllTiles(f func(t *Tile)) {
	for _, c := range ch.tiles {
		f(c)
	}
}

func (ch *chunk) Tile(tileCoord Coord) *Tile {
	return ch.tiles[ch.tileIndex(tileCoord.X(), tileCoord.Y())]

}

func (ch *chunk) initTile(c Coord) *Tile {
	return &Tile{c, make(map[int]int)}
}

func (ch *chunk) tileIndex(x, y int) int {
	x -= ch.location.X() * ChunkSize
	y -= ch.location.Y() * ChunkSize
	return (x * ChunkSize) + y
}
