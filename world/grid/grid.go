// A package for managing a grid based game world
package grid

import (
	"github.com/pedritoelcabra/projectx/core/logger"
	"github.com/pedritoelcabra/projectx/world/coord"
	"github.com/pedritoelcabra/projectx/world/defs"
	"github.com/pedritoelcabra/projectx/world/noise"
	"github.com/pedritoelcabra/projectx/world/tiling"
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

func (g *Grid) CreateNewChunk(chunkCoord coord.Coord) {
	chunkIndex := g.chunkIndex(chunkCoord.X(), chunkCoord.Y())
	aChunk := NewChunk(chunkCoord)
	g.Chunks[chunkIndex] = aChunk
	g.chunksToGenerate = append(g.chunksToGenerate, chunkCoord)
	aChunk.queuedForGeneration = true
	aChunk.Generated = false
}

func (g *Grid) ChunkGeneration(playerTile coord.Coord, tick int) {
	g.GenerateChunk()
	if tick%60 > 0 {
		return
	}
	playerChunk := g.ChunkCoord(playerTile)
	for x := playerChunk.X() - 3; x <= playerChunk.X()+3; x++ {
		for y := playerChunk.Y() - 3; y <= playerChunk.Y()+3; y++ {
			chunkIndex := g.chunkIndex(x, y)
			chunkCoord := coord.NewCoord(x, y)
			aChunk, chunkExists := g.Chunks[chunkIndex]
			if !chunkExists {
				g.CreateNewChunk(chunkCoord)
				continue
			}
			if !aChunk.Generated && !aChunk.queuedForGeneration {
				g.QueueChunkForGeneration(chunkCoord)
			}
		}
	}
}

func (g *Grid) QueueChunkForGeneration(chunkCoord coord.Coord) {
	g.chunksToGenerate = append(g.chunksToGenerate, chunkCoord)
}

func (g *Grid) GenerateChunk() {
	if len(g.chunksToGenerate) < 1 {
		return
	}
	chunkCoord := g.chunksToGenerate[0]
	g.chunksToGenerate = g.chunksToGenerate[1:]
	aChunk := g.Chunk(chunkCoord)
	aChunk.queuedForGeneration = false
	aChunk.Generated = true
	logger.General("Generated chunk: "+chunkCoord.ToString(), nil)
}

func (t *Tile) InitializeTile() {
	height := noise.Generator.GetHeight(t.X(), t.Y())
	t.Set(Height, height)
	terrain := -1
	terrain = defs.BasicMountain
	if height < 600 {
		terrain = defs.BasicDesert
	}
	if height < 200 {
		terrain = defs.BasicGrass
	}
	if height < 0 {
		terrain = defs.BasicWater
	}
	t.Set(TerrainBase, terrain)
	pixelX, pixelY := tiling.TileIToPixelF(t.X(), t.Y())
	t.SetF(RenderX, pixelX)
	t.SetF(RenderY, pixelY)
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
