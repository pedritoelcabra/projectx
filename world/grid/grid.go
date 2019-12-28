// A package for managing a grid based game world
package grid

import (
	"github.com/pedritoelcabra/projectx/world/coord"
	"github.com/pedritoelcabra/projectx/world/defs"
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

type ChunkMap map[int]*chunk

type Grid struct {
	Chunks           ChunkMap
	chunksToGenerate []coord.Coord
	noise            *noise.NoiseGenerator
}

func New() *Grid {
	arraySize := GridSize * GridSize
	arrayChunks := make(map[int]*chunk, arraySize)
	aGrid := &Grid{}
	aGrid.Chunks = arrayChunks
	return aGrid
}

func (g *Grid) SetNoise(noise *noise.NoiseGenerator) {
	g.noise = noise
}

func (g *Grid) Tile(tileCoord coord.Coord) *Tile {
	return g.Chunk(g.ChunkCoord(tileCoord)).Tile(tileCoord)
}

func (g *Grid) Chunk(chunkCoord coord.Coord) *chunk {
	chunkIndex := g.chunkIndex(chunkCoord.X(), chunkCoord.Y())
	if aChunk, chunkExists := g.Chunks[chunkIndex]; chunkExists {
		if !aChunk.isPreloaded {
			aChunk.Preload(chunkCoord)
		}
		return aChunk
	}
	g.CreateNewChunk(chunkCoord)
	return g.Chunks[chunkIndex]
}

func (t *Tile) InitializeTile() {
	height := noise.Generator.GetHeight(t.X(), t.Y())
	t.Set(Height, height)
	terrain := -1
	terrain = defs.BasicMountain
	if height < 300 {
		terrain = defs.BasicHills
	}
	if height < 150 {
		terrain = defs.BasicGrass
	}
	if height < 0 {
		terrain = defs.BasicWater
	}
	if height < -50 {
		terrain = defs.BasicDeepWater
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

func (g *Grid) ChunkCoord(tileCoord coord.Coord) coord.Coord {
	x := ((tileCoord.X() + TileOffset) / ChunkSize) - GridOffset
	y := ((tileCoord.Y() + TileOffset) / ChunkSize) - GridOffset
	return coord.NewCoord(x, y)
}
