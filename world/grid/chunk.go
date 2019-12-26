package grid

import (
	"github.com/pedritoelcabra/projectx/core/logger"
	"github.com/pedritoelcabra/projectx/world/coord"
	"github.com/pedritoelcabra/projectx/world/tiling"
)

type chunk struct {
	tiles               []*Tile
	Location            coord.Coord
	Generated           bool
	queuedForGeneration bool
	isPreloaded         bool
}

func NewChunk(location coord.Coord) *chunk {
	aChunk := &chunk{}
	aChunk.isPreloaded = false
	aChunk.Preload(location)
	return aChunk
}

func (ch *chunk) Preload(location coord.Coord) {
	if ch.isPreloaded {
		return
	}
	ch.tiles = make([]*Tile, ChunkSize*ChunkSize)
	ch.Location = location
	for x := 0; x < ChunkSize; x++ {
		for y := 0; y < ChunkSize; y++ {
			tileX := (ch.Location.X() * ChunkSize) + x
			tileY := (ch.Location.Y() * ChunkSize) + y
			tileLocation := coord.NewCoord(tileX, tileY)
			tileIndex := ch.tileIndex(tileX, tileY)
			aTile := &Tile{}
			aTile.values = make(map[int]int)
			aTile.valuesF = make(map[int]float64)
			aTile.coordinates = tileLocation

			centerX := float64(tileX) * tiling.TileHorizontalSeparation
			centerY := float64(tileY) * tiling.TileHeight
			if x%2 > 0 {
				centerY += tiling.TileHeight / 2
			}
			renderX := centerX - tiling.TileWidth/2
			renderY := centerY - tiling.TileHeight/2
			aTile.SetF(RenderX, renderX)
			aTile.SetF(RenderY, renderY)
			aTile.SetF(CenterX, centerX)
			aTile.SetF(CenterY, centerY)
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

func (ch *chunk) Tile(tileCoord coord.Coord) *Tile {
	return ch.tiles[ch.tileIndex(tileCoord.X(), tileCoord.Y())]
}

func (ch *chunk) tileIndex(x, y int) int {
	x -= ch.Location.X() * ChunkSize
	y -= ch.Location.Y() * ChunkSize
	return (x * ChunkSize) + y
}
