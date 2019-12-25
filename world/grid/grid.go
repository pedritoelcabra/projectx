// A package for managing a grid based game world
package grid

import (
	"github.com/pedritoelcabra/projectx/core/logger"
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
	chunksToGenerate []Coord
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

func (g *Grid) Tile(tileCoord Coord) *Tile {
	return g.Chunk(g.chunkCoord(tileCoord)).Tile(tileCoord)
}

func (g *Grid) Chunk(chunkCoord Coord) *chunk {
	chunkIndex := g.chunkIndex(chunkCoord.X(), chunkCoord.Y())
	if aChunk, chunkExists := g.Chunks[chunkIndex]; chunkExists {
		return aChunk
	}
	g.PreLoadChunk(chunkCoord)
	return g.Chunks[chunkIndex]
}

func (g *Grid) PreLoadChunk(chunkCoord Coord) {
	chunkIndex := g.chunkIndex(chunkCoord.X(), chunkCoord.Y())
	aChunk := NewChunk(chunkCoord)
	aChunk.RunOnAllTiles(func(t *Tile) {
		t.InitializeTile(g)
	})
	g.Chunks[chunkIndex] = aChunk
	g.chunksToGenerate = append(g.chunksToGenerate, chunkCoord)
	aChunk.queuedForGeneration = true
	aChunk.Generated = false
	logger.General("Preloaded chunk: "+chunkCoord.ToString(), nil)
}

func (g *Grid) ChunkGeneration(playerTile Coord, tick int) {
	g.GenerateChunk()
	if tick%60 > 0 {
		return
	}
	playerChunk := g.chunkCoord(playerTile)
	for x := playerChunk.X() - 3; x <= playerChunk.X()+3; x++ {
		for y := playerChunk.Y() - 3; y <= playerChunk.Y()+3; y++ {
			chunkIndex := g.chunkIndex(x, y)
			chunkCoord := NewCoord(x, y)
			aChunk, chunkExists := g.Chunks[chunkIndex]
			if !chunkExists {
				g.PreLoadChunk(chunkCoord)
				continue
			}
			if !aChunk.Generated && !aChunk.queuedForGeneration {
				g.chunksToGenerate = append(g.chunksToGenerate, chunkCoord)
			}
		}
	}
}

func (g *Grid) GenerateChunk() {
	if len(g.chunksToGenerate) < 1 {
		return
	}
	chunkCoord := g.chunksToGenerate[0]
	g.chunksToGenerate = g.chunksToGenerate[1:]
	logger.General("Generateed chunk: "+chunkCoord.ToString(), nil)
	aChunk := g.Chunk(chunkCoord)
	aChunk.queuedForGeneration = false
	aChunk.Generated = true
}

func (t *Tile) InitializeTile(g *Grid) {
	height := g.noise.GetHeight(t.X(), t.Y())
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
