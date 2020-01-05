// A package for managing a grid based game world
package world

import (
	tiling2 "github.com/pedritoelcabra/projectx/src/world/tiling"
	utils2 "github.com/pedritoelcabra/projectx/src/world/utils"
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
	chunksToGenerate []tiling2.Coord
	noise            *utils2.NoiseGenerator
}

func New() *Grid {
	arraySize := GridSize * GridSize
	arrayChunks := make(map[int]*chunk, arraySize)
	aGrid := &Grid{}
	aGrid.Chunks = arrayChunks
	return aGrid
}

func (g *Grid) SetNoise(noise *utils2.NoiseGenerator) {
	g.noise = noise
}

func (g *Grid) Tile(tileCoord tiling2.Coord) *Tile {
	return g.Chunk(g.ChunkCoord(tileCoord)).Tile(tileCoord)
}

func (g *Grid) Chunk(chunkCoord tiling2.Coord) *chunk {
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

func (g *Grid) chunkIndex(x, y int) int {
	x += GridOffset
	y += GridOffset
	if x < 0 || x >= GridSize || y < 0 || y >= GridSize {
		log.Fatalf("Grid.Tile() requested invalid chunk %d / %d", x, y)
	}
	return (x * GridSize) + y
}

func (g *Grid) ChunkCoord(tileCoord tiling2.Coord) tiling2.Coord {
	x := ((tileCoord.X() + TileOffset) / ChunkSize) - GridOffset
	y := ((tileCoord.Y() + TileOffset) / ChunkSize) - GridOffset
	return tiling2.NewCoord(x, y)
}
