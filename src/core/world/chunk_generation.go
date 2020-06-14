package world

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	tiling2 "github.com/pedritoelcabra/projectx/src/core/world/tiling"
	utils2 "github.com/pedritoelcabra/projectx/src/core/world/utils"
)

func (g *Grid) CreateNewChunk(chunkCoord tiling2.Coord) {
	chunkIndex := g.chunkIndex(chunkCoord.X(), chunkCoord.Y())
	aChunk := NewChunk(chunkCoord)
	g.Chunks[chunkIndex] = aChunk
	aChunk.Generated = false
}

func (g *Grid) ChunkGeneration() {
	g.ProcessChunkGenerationQueue()
	if theWorld.GetTick() > 100 && !theWorld.IsTock() {
		return
	}
	for _, chunk := range ChunksAroundPlayer(3) {
		if !chunk.Generated && !chunk.queuedForGeneration {
			g.QueueChunkForGeneration(chunk)
		}
	}
}

func (g *Grid) QueueChunkForGeneration(aChunk *Chunk) {
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
	aChunk.RunOnAllTiles(func(t *Tile) {
		t.GenerateResources()
	})
	g.SpawnSector(aChunk)
	aChunk.RunOnAllTiles(func(t *Tile) {
		t.Recalculate()
	})
	aChunk.queuedForGeneration = false
	aChunk.Generated = true
	//logger.General("Generated Chunk: "+chunkCoord.ToString(), nil)
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

func (t *Tile) GenerateResources() {
	if t.IsImpassable() || t.Get(Height) <= 0 {
		return
	}

	x := t.X()
	y := t.Y()

	ironScore := utils2.Generator.GetIron(x, y)
	if ironScore > 350 {
		t.SetResourceByName("Iron Ore")
		return
	}

	coalScore := utils2.Generator.GetCoal(x, y)
	if coalScore > 350 {
		t.SetResourceByName("Coal")
		return
	}

	stoneScore := utils2.Generator.GetStone(x, y)
	if stoneScore > 350 {
		t.SetResourceByName("Stone")
		return
	}

	bioMassScore := utils2.Generator.GetBiomass(x, y)
	bioMassScore += randomizer.RandomInt(0, 500)
	if bioMassScore > 200 {
		t.SetResourceByName("Deciduous Forest Sparse")
		return
	}
	if bioMassScore > 350 {
		t.SetResourceByName("Deciduous Forest")
		return
	}
	t.SetResourceByName("")
}

func (t *Tile) SetResourceByName(resourceName string) {
	if resourceName == "" {
		return
	}

	resourceId := defs.ResourceByName(resourceName)
	resourceDef := defs.ResourceById(resourceId)
	defs.AddResourceLocation(resourceDef.Resource)
	t.Set(Resource, resourceId)
	t.Set(ResourceAmount, resourceDef.ResourceAmount)
}
