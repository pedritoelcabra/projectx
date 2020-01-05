package world

import (
	tiling2 "github.com/pedritoelcabra/projectx/src/world/tiling"
	utils2 "github.com/pedritoelcabra/projectx/src/world/utils"
)

func (g *Grid) CreateNewChunk(chunkCoord tiling2.Coord) {
	chunkIndex := g.chunkIndex(chunkCoord.X(), chunkCoord.Y())
	aChunk := NewChunk(chunkCoord)
	g.Chunks[chunkIndex] = aChunk
	aChunk.Generated = false
}

func (g *Grid) ChunkGeneration(playerTile tiling2.Coord, tick int) {
	g.ProcessChunkGenerationQueue()
	if tick > 100 && tick%60 > 0 {
		return
	}
	playerChunk := g.ChunkCoord(playerTile)
	for x := playerChunk.X() - 3; x <= playerChunk.X()+3; x++ {
		for y := playerChunk.Y() - 3; y <= playerChunk.Y()+3; y++ {
			chunkIndex := g.chunkIndex(x, y)
			chunkCoord := tiling2.NewCoord(x, y)
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

func (g *Grid) GenerateChunk(chunkCoord tiling2.Coord) {
	aChunk := g.Chunk(chunkCoord)
	if aChunk.IsGenerated() {
		return
	}
	g.SpawnSector(aChunk)
	aChunk.queuedForGeneration = false
	aChunk.Generated = true
	//logger.General("Generated chunk: "+chunkCoord.ToString(), nil)
}

func (t *Tile) InitializeTile() {
	height := utils2.Generator.GetHeight(t.X(), t.Y())
	biomeValue := utils2.Generator.GetBiome(t.X(), t.Y())
	biomeValue = 0
	biome := utils2.BiomeTemperate
	if biomeValue > 250 {
		biome = utils2.BiomeDesert
	}
	if biomeValue < -250 {
		biome = utils2.BiomeTundra
	}
	t.Set(Biome, biome)
	t.Set(Height, height)
	t.Set(SectorId, -1)
	t.SetTerrain()
	t.Recalculate()
}

func (t *Tile) SetTerrain() {
	height := t.Get(Height)
	terrain := -1
	terrain = utils2.BasicMountain
	if height < 300 {
		terrain = utils2.BasicHills
	}
	if height < 150 {
		terrain = utils2.BasicGrass
	}
	if height < 0 {
		terrain = utils2.BasicWater
	}
	if height < -50 {
		terrain = utils2.BasicDeepWater
	}
	if t.Get(Biome) == utils2.BiomeTundra {
		terrain += 10
	}
	if t.Get(Biome) == utils2.BiomeDesert {
		terrain += 20
	}
	t.Set(TerrainBase, terrain)
}
