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
	minX := aChunk.FirstTile().X() + 5
	maxX := aChunk.FirstTile().X() + ChunkSize - 5
	minY := aChunk.FirstTile().Y() + 5
	maxY := aChunk.FirstTile().Y() + ChunkSize - 5
	for attempts := 0; attempts <= 10; attempts++ {
		randomX := randomizer.RandomInt(minX, maxX)
		randomY := randomizer.RandomInt(minY, maxY)
		aTile := g.Tile(coord.NewCoord(randomX, randomY))
		if g.TileIsSuitableForSectorCenter(aTile) {
			return aTile.coordinates, true
		}
	}
	logger.General("Failed to find a suitable center for sector in chunk: "+aChunk.Location.ToString(), nil)
	return aChunk.FirstTile().coordinates, false
}

func (g *Grid) TileIsSuitableForSectorCenter(aTile *Tile) bool {
	necessarySpace := 3
	for x := aTile.X() - necessarySpace; x <= aTile.X()+necessarySpace; x++ {
		for y := aTile.Y() - necessarySpace; y <= aTile.Y()+necessarySpace; y++ {
			nearbyTile := g.Tile(coord.NewCoord(x, y))
			if nearbyTile.IsImpassable() {
				return false
			}
		}
	}
	return true
}

func (g *Grid) ShouldSpawnSector(aChunk *chunk) bool {
	nearbySectorCount := 0
	if aChunk.ChunkData.Get(AvgHeight) < 0 {
		return false
	}
	radiusToCheck := 3
	for x := aChunk.Location.X() - radiusToCheck; x <= aChunk.Location.X()+radiusToCheck; x++ {
		for y := aChunk.Location.Y() - radiusToCheck; y <= aChunk.Location.Y()+radiusToCheck; y++ {
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

	// 80% chance to generate sector at no neighbours, 15% less for each nearby sector
	chanceToGenerateSector := 80
	chanceToGenerateSector -= nearbySectorCount * 15
	if !randomizer.PercentageRoll(chanceToGenerateSector) {
		return false
	}
	return true
}
