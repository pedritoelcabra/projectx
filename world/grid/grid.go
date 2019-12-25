// A package for managing a grid based game world
package grid

import (
	"github.com/pedritoelcabra/projectx/core/logger"
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/world/noise"
	"log"
)

const (
	ChunkSize  = 32
	GridSize   = 1000
	GridOffset = 500
	GridTiles  = 32000
	TileOffset = 16000
)

type Grid struct {
	chunks map[int]*chunk
	noise  *noise.NoiseGenerator
}

func New() *Grid {
	arraySize := GridSize * GridSize
	arrayChunks := make(map[int]*chunk, arraySize)
	aGrid := &Grid{}
	aGrid.chunks = arrayChunks
	return aGrid
}

func (g *Grid) SetNoise(noise *noise.NoiseGenerator) {
	g.noise = noise
}

func (g *Grid) Tile(tileCoord Coord) *Tile {
	chunkCoord := g.chunkCoord(tileCoord)
	chunkIndex := g.chunkIndex(chunkCoord.X(), chunkCoord.Y())
	if aChunk, chunkExists := g.chunks[chunkIndex]; chunkExists {
		return aChunk.Tile(tileCoord)
	}
	g.PreLoadChunk(chunkCoord)
	return g.chunks[chunkIndex].Tile(tileCoord)
}

func (g *Grid) PreLoadChunk(chunkCoord Coord) {
	chunkIndex := g.chunkIndex(chunkCoord.X(), chunkCoord.Y())
	aChunk := NewChunk(chunkCoord)
	aChunk.RunOnAllTiles(func(t *Tile) {
		t.InitializeTile(g)
	})
	g.chunks[chunkIndex] = aChunk
	logger.General("Preloaded chunk: "+chunkCoord.ToString(), nil)
}

func (t *Tile) InitializeTile(g *Grid) {
	height := g.noise.GetHeight(t.X(), t.Y())
	t.Set(Height, height)
	terrain := -1
	terrain = gfx.StoneBlockFull
	if height < 600 {
		terrain = gfx.DesertFull
	}
	if height < 200 {
		terrain = gfx.GrassFull
	}
	if height < 0 {
		terrain = gfx.WaterFull
	}
	t.Set(TerrainBase, terrain)
}

func (g *Grid) chunkIndex(x, y int) int {
	x += GridOffset
	y += GridOffset
	if x < 0 || x >= GridSize || y < 0 || y >= GridSize {
		log.Fatalf("Grid.Tile() requested invalid chunk %d / %d", x, y)
	}
	return (x * GridSize) + y
}

func (g *Grid) chunkCoord(tileCoord Coord) Coord {
	x := ((tileCoord.X() + TileOffset) / ChunkSize) - GridOffset
	y := ((tileCoord.Y() + TileOffset) / ChunkSize) - GridOffset
	return NewCoord(x, y)
}
