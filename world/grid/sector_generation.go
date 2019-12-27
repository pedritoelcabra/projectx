package grid

import (
	"github.com/pedritoelcabra/projectx/core/logger"
	"github.com/pedritoelcabra/projectx/core/randomizer"
	"github.com/pedritoelcabra/projectx/world/coord"
)

func (g *Grid) SpawnSector(aChunk *chunk) {
	if !g.ShouldSpawnSector(aChunk) {
		return
	}
	centerCoord, hasValidCenter := g.SuitableSectorCenter(aChunk)
	if !hasValidCenter {
		return
	}
	aChunk.Sector = NewSector(centerCoord)
	logger.General("Spawned sector in chunk: "+aChunk.Location.ToString(), nil)
}

func (g *Grid) SuitableSectorCenter(aChunk *chunk) (coord.Coord, bool) {
	return aChunk.FirstTile().coordinates, true
}

func (g *Grid) ShouldSpawnSector(aChunk *chunk) bool {
	nearbySectorCount := 0
	if aChunk.ChunkData.Get(AvgHeight) < 0 {
		return false
	}
	for x := aChunk.Location.X() - 2; x <= aChunk.Location.X()+2; x++ {
		for y := aChunk.Location.Y() - 2; y <= aChunk.Location.Y()+2; y++ {
			bChunk := g.Chunk(coord.NewCoord(x, y))
			if bChunk.Sector == nil {
				continue
			}
			if aChunk.Location.ChebyshevDist(bChunk.Location) <= 1 {
				return false
			}
			nearbySectorCount++
		}
	}

	// 80% chance to generate sector at no neighbours, 20% with the maximum (4) neighbours in the other circle
	chanceToGenerateSector := 80
	chanceToGenerateSector -= nearbySectorCount * 15
	if !randomizer.PercentageRoll(chanceToGenerateSector) {
		return false
	}
	return true
}
