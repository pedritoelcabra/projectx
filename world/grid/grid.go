// A package for managing a grid based game world
package grid

import (
	"github.com/pedritoelcabra/projectx/world/tiling"
	"github.com/pedritoelcabra/projectx/world/utils"
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
	chunksToGenerate []tiling.Coord
	noise            *utils.NoiseGenerator
}

func New() *Grid {
	arraySize := GridSize * GridSize
	arrayChunks := make(map[int]*chunk, arraySize)
	aGrid := &Grid{}
	aGrid.Chunks = arrayChunks
	return aGrid
}

func (g *Grid) SetNoise(noise *utils.NoiseGenerator) {
	g.noise = noise
}

func (g *Grid) Tile(tileCoord tiling.Coord) *Tile {
	return g.Chunk(g.ChunkCoord(tileCoord)).Tile(tileCoord)
}

func (g *Grid) Chunk(chunkCoord tiling.Coord) *chunk {
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
	height := utils.Generator.GetHeight(t.X(), t.Y())
	t.Set(Height, height)
	terrain := -1
	terrain = utils.BasicMountain
	if height < 300 {
		terrain = utils.BasicHills
	}
	if height < 150 {
		terrain = utils.BasicGrass
	}
	if height < 0 {
		terrain = utils.BasicWater
	}
	if height < -50 {
		terrain = utils.BasicDeepWater
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

func (g *Grid) ChunkCoord(tileCoord tiling.Coord) tiling.Coord {
	x := ((tileCoord.X() + TileOffset) / ChunkSize) - GridOffset
	y := ((tileCoord.Y() + TileOffset) / ChunkSize) - GridOffset
	return tiling.NewCoord(x, y)
}
