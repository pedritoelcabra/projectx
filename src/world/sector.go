package world

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/world/container"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"log"
	"math"
	"strconv"
)

type SectorKey int
type SectorMap map[SectorKey]*Sector

type Sector struct {
	Id       SectorKey
	Size     int
	Name     string
	Center   tiling.Coord
	Data     *container.Container
	Template *defs.SectorDef
	Tiles    []tiling.Coord
}

func NewSector(location tiling.Coord, def *defs.SectorDef) *Sector {
	aSector := &Sector{}
	aSector.Template = def
	aSector.Data = container.NewContainer()
	aSector.Center = location
	aSector.Id = theWorld.AddSector(aSector)
	aSector.Name = aSector.Template.Name + " " + strconv.Itoa(int(aSector.Id))
	aSector.AddTile(aSector.Center)
	aSector.GrowSectorToSize(def.Size, aSector.Center)
	aSector.SpawnBuildings()
	aSector.Init()
	return aSector
}

func (s *Sector) SpawnBuildings() {
	NewBuilding(s.Template.CenterBuilding, theWorld.Grid.Tile(s.Center))
	for name, chance := range s.Template.Buildings {
		firstSpawned := false
		for remainingChance := chance; remainingChance > 0; {
			spawn := false
			def := defs.GetBuildingDef(name)
			if def == nil {
				log.Fatal("Invalid building name: " + name)
			}
			if !firstSpawned && remainingChance >= 100 {
				remainingChance -= 100
				spawn = true
				firstSpawned = true
			}
			if !spawn {
				thisChance := 50
				if remainingChance < 50 {
					thisChance = remainingChance
				}
				remainingChance -= thisChance
				spawn = randomizer.PercentageRoll(thisChance)
			}
			if spawn {
				s.AttemptSpawnBuilding(def)
			}
		}
	}
}

func (s *Sector) AttemptSpawnBuilding(def *defs.BuildingDef) {
	bestScore := -1000
	bestLocation := tiling.NewCoord(0, 0)
	for _, tileCoord := range s.Tiles {
		tile := theWorld.Grid.Tile(tileCoord)
		if tile.IsImpassable() || !tile.IsLand() {
			continue
		}
		if tile.GetBuilding() != nil {
			continue
		}
		score := randomizer.RandomInt(0, 20)
		score -= tile.GetCoord().ChebyshevDist(s.GetCenter())
		if score > bestScore {
			bestScore = score
			bestLocation = tileCoord
		}
	}
	if bestScore > -1000 {
		NewBuilding(def.Name, theWorld.Grid.Tile(bestLocation))
	}
}

func (s *Sector) GrowSectorToSize(size int, centerCoord tiling.Coord) {
	s.Size = size
	options := NewPathOptions()
	options.MinMoveCost = 1.0
	sizeF := float64(size)
	for x := centerCoord.X() - size; x <= centerCoord.X()+size; x++ {
		for y := centerCoord.Y() - size; y <= centerCoord.Y()+size; y++ {
			aCoord := tiling.NewCoord(x, y)
			if theWorld.Grid.Tile(aCoord).Get(SectorId) >= 0 {
				continue
			}
			path := FindPathWithOptions(centerCoord, aCoord, options)
			if path.IsValid() && math.Floor(path.GetCost()) <= sizeF {
				s.AddTile(aCoord)
			}
		}
	}
	s.RecalculateTiles()
}

func (s *Sector) AddTile(tileCoord tiling.Coord) {
	for _, existantTiles := range s.Tiles {
		if existantTiles.Equals(tileCoord) {
			return
		}
	}
	s.Tiles = append(s.Tiles, tileCoord)
	theWorld.Grid.Tile(tileCoord).Set(SectorId, int(s.Id))
}

func (s *Sector) RecalculateTiles() {
	for _, tile := range s.Tiles {
		theWorld.Grid.Tile(tile).Set(SectorId, int(s.Id))
	}
	for _, tile := range s.Tiles {
		theWorld.Grid.Tile(tile).Recalculate()
	}
}

func (s *Sector) GetCenter() tiling.Coord {
	return s.Center
}

func (s *Sector) GetName() string {
	return s.Name
}

func (s *Sector) GetId() SectorKey {
	return s.Id
}

func (s *Sector) Init() {
	theWorld.Grid.Chunk(theWorld.Grid.ChunkCoord(s.Center)).SetSector(s)
	s.RecalculateTiles()
}

func (s *Sector) Get(key int) int {
	return s.Data.Get(key)
}

func (s *Sector) GetF(key int) float64 {
	return s.Data.GetF(key)
}

func (s *Sector) Set(key, value int) {
	s.Data.Set(key, value)
}

func (s *Sector) SetF(key int, value float64) {
	s.Data.SetF(key, value)
}

func (s *Sector) SetName(name string) {
	s.Name = name
}
