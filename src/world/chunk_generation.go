package world

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"github.com/pedritoelcabra/projectx/src/world/utils"
)

func (g *Grid) CreateNewChunk(chunkCoord tiling.Coord) {
	chunkIndex := g.chunkIndex(chunkCoord.X(), chunkCoord.Y())
	aChunk := NewChunk(chunkCoord)
	g.Chunks[chunkIndex] = aChunk
	aChunk.Generated = false
}

func (g *Grid) ChunkGeneration(playerTile tiling.Coord, tick int) {
	g.ProcessChunkGenerationQueue()
	if tick > 100 && tick%60 > 0 {
		return
	}
	playerChunk := g.ChunkCoord(playerTile)
	for x := playerChunk.X() - 3; x <= playerChunk.X()+3; x++ {
		for y := playerChunk.Y() - 3; y <= playerChunk.Y()+3; y++ {
			chunkIndex := g.chunkIndex(x, y)
			chunkCoord := tiling.NewCoord(x, y)
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

func (g *Grid) GenerateChunk(chunkCoord tiling.Coord) {
	aChunk := g.Chunk(chunkCoord)
	if aChunk.IsGenerated() {
		return
	}
	aChunk.RunOnAllTiles(func(t *Tile) {
		t.GenerateVegetation()
	})
	g.SpawnSector(aChunk)
	aChunk.RunOnAllTiles(func(t *Tile) {
		t.Recalculate()
	})
	aChunk.queuedForGeneration = false
	aChunk.Generated = true
	//logger.General("Generated chunk: "+chunkCoord.ToString(), nil)
}

func (t *Tile) InitializeTile() {
	height := utils.Generator.GetHeight(t.X(), t.Y())
	biomeValue := utils.Generator.GetBiome(t.X(), t.Y())
	biomeValue = 0
	biome := utils.BiomeTemperate
	if biomeValue > 250 {
		biome = utils.BiomeDesert
	}
	if biomeValue < -250 {
		biome = utils.BiomeTundra
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
	if t.Get(Biome) == utils.BiomeTundra {
		terrain += 10
	}
	if t.Get(Biome) == utils.BiomeDesert {
		terrain += 20
	}
	t.Set(TerrainBase, terrain)
}

func (t *Tile) GenerateVegetation() {
	if t.IsImpassable() || t.Get(Height) <= 0 {
		return
	}
	bioMassScore := utils.Generator.GetBiomass(t.X(), t.Y())
	bioMassScore += randomizer.RandomInt(0, 500)
	vegName := ""
	if bioMassScore > 200 {
		vegName = "Deciduous Forest Sparse"
	}
	if bioMassScore > 350 {
		vegName = "Deciduous Forest"
	}
	if vegName == "" {
		return
	}
	t.Set(Flora, defs.VegetationByName(vegName))
}
