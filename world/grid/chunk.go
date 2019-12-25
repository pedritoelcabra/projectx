package grid

import "github.com/pedritoelcabra/projectx/core/logger"

type chunk struct {
	tiles               []*Tile
	Location            Coord
	Generated           bool
	queuedForGeneration bool
	isPreloaded         bool
}

func NewChunk(location Coord) *chunk {
	aChunk := &chunk{}
	aChunk.isPreloaded = false
	aChunk.Preload(location)
	return aChunk
}

func (ch *chunk) Preload(location Coord) {
	if ch.isPreloaded {
		return
	}
	ch.tiles = make([]*Tile, ChunkSize*ChunkSize)
	ch.Location = location
	for x := 0; x < ChunkSize; x++ {
		for y := 0; y < ChunkSize; y++ {
			tileX := (ch.Location.X() * ChunkSize) + x
			tileY := (ch.Location.Y() * ChunkSize) + y
			tileLocation := NewCoord(tileX, tileY)
			tileIndex := ch.tileIndex(tileX, tileY)
			aTile := &Tile{tileLocation, make(map[int]int)}
			ch.tiles[tileIndex] = aTile
		}
	}
	ch.RunOnAllTiles(func(t *Tile) {
		t.InitializeTile()
	})
	ch.isPreloaded = true
	logger.General("Preloaded chunk: "+location.ToString(), nil)
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
	x -= ch.Location.X() * ChunkSize
	y -= ch.Location.Y() * ChunkSize
	return (x * ChunkSize) + y
}
