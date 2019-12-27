package grid

import (
	"github.com/pedritoelcabra/projectx/world/coord"
)

func (g *Grid) CreateNewChunk(chunkCoord coord.Coord) {
	chunkIndex := g.chunkIndex(chunkCoord.X(), chunkCoord.Y())
	aChunk := NewChunk(chunkCoord)
	g.Chunks[chunkIndex] = aChunk
	aChunk.Generated = false
}

func (g *Grid) ChunkGeneration(playerTile coord.Coord, tick int) {
	g.ProcessChunkGenerationQueue()
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
				g.QueueChunkForGeneration(aChunk)
			}
		}
	}
}

func (g *Grid) QueueChunkForGeneration(aChunk *chunk) {
	g.chunksToGenerate = append(g.chunksToGenerate, aChunk.Location)
	aChunk.queuedForGeneration = true
}

func (g *Grid) ProcessChunkGenerationQueue() {
	if len(g.chunksToGenerate) < 1 {
		return
	}
	chunkCoord := g.chunksToGenerate[0]
	g.chunksToGenerate = g.chunksToGenerate[1:]
	g.GenerateChunk(chunkCoord)
}

func (g *Grid) GenerateChunk(chunkCoord coord.Coord) {
	aChunk := g.Chunk(chunkCoord)
	if aChunk.IsGenerated() {
		return
	}
	g.SpawnSector(aChunk)
	aChunk.queuedForGeneration = false
	aChunk.Generated = true
	//logger.General("Generated chunk: "+chunkCoord.ToString(), nil)
}
