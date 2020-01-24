package world

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/logger"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"strconv"
)

func (g *Grid) SpawnSector(aChunk *chunk) {
	if !g.ShouldSpawnSector(aChunk) {
		return
	}
	centerCoord, hasValidCenter := g.SuitableSectorCenter(aChunk)
	if !hasValidCenter {
		return
	}
	template := g.ChooseSectorTemplate(centerCoord)
	if template == nil {
		return
	}
	newSector := NewSector(centerCoord, template)
	newFaction := NewFaction("Faction " + strconv.Itoa(len(theWorld.Entities.Factions)))
	newSector.Set(FactionId, int(newFaction.GetId()))
	tile := g.Tile(centerCoord)
	tile.Set(Flora, 0)
	//logger.General("Spawned a "+template.Name+" sector in chunk: "+aChunk.Location.ToString()+" at "+centerCoord.ToString(), nil)
}

func (g *Grid) ChooseSectorTemplate(location tiling.Coord) *defs.SectorDef {
	sectorDefs := defs.SectorDefs()
	bestDef := &defs.SectorDef{}
	bestDef = nil
	bestScore := -1
	for _, sectorDef := range sectorDefs {
		score := sectorDef.Weight
		score *= randomizer.RandomInt(0, 100)
		if theWorld.GetTick() == 0 && sectorDef.Name == "Player Village" {
			return sectorDef
		}
		if score > bestScore {
			bestDef = sectorDef
			bestScore = score
		}
	}
	return bestDef
}

func (g *Grid) SuitableSectorCenter(aChunk *chunk) (tiling.Coord, bool) {
	minX := aChunk.FirstTile().X() + 5
	maxX := aChunk.FirstTile().X() + ChunkSize - 5
	minY := aChunk.FirstTile().Y() + 5
	maxY := aChunk.FirstTile().Y() + ChunkSize - 5
	for attempts := 0; attempts <= 10; attempts++ {
		randomX := randomizer.RandomInt(minX, maxX)
		randomY := randomizer.RandomInt(minY, maxY)
		aTile := g.Tile(tiling.NewCoord(randomX, randomY))
		if g.TileIsSuitableForSectorCenter(aTile) {
			return aTile.coordinates, true
		}
	}
	logger.General("Failed to find a suitable center for sector in chunk: "+aChunk.Location.ToString(), nil)
	return aChunk.FirstTile().coordinates, false
}

func (g *Grid) TileIsSuitableForSectorCenter(aTile *Tile) bool {
	necessarySpace := 3
	maxImpassableTiles := 5
	currentImpassableTiles := 0
	for x := aTile.X() - necessarySpace; x <= aTile.X()+necessarySpace; x++ {
		for y := aTile.Y() - necessarySpace; y <= aTile.Y()+necessarySpace; y++ {
			nearbyTile := g.Tile(tiling.NewCoord(x, y))
			if nearbyTile.IsImpassable() || !nearbyTile.IsLand() {
				currentImpassableTiles++
			}
		}
	}
	return currentImpassableTiles <= maxImpassableTiles
}

func (g *Grid) ShouldSpawnSector(aChunk *chunk) bool {
	nearbySectorCount := 0
	if aChunk.ChunkData.Get(AvgHeight) < 0 {
		return false
	}
	radiusToCheck := 3
	for x := aChunk.Location.X() - radiusToCheck; x <= aChunk.Location.X()+radiusToCheck; x++ {
		for y := aChunk.Location.Y() - radiusToCheck; y <= aChunk.Location.Y()+radiusToCheck; y++ {
			bChunk := g.Chunk(tiling.NewCoord(x, y))
			if bChunk.GetSector() == nil {
				continue
			}
			if aChunk.Location.ChebyshevDist(bChunk.Location) <= 1 {
				return false
			}
			nearbySectorCount++
		}
	}

	// 80% chance to generate sector at no neighbours, 15% less for each nearby sector
	chanceToGenerateSector := 80
	chanceToGenerateSector -= nearbySectorCount * 15
	if !randomizer.PercentageRoll(chanceToGenerateSector) {
		return false
	}
	return true
}
