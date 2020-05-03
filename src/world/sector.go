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
type ConnectionMap map[SectorKey]SectorConnection
type BuildingList []BuildingPointer

type Sector struct {
	Id            SectorKey
	Size          int
	Name          string
	Center        tiling.Coord
	Data          *container.Container
	Template      *defs.SectorDef
	Tiles         []tiling.Coord
	NearbySectors ConnectionMap
}

func NewSector(location tiling.Coord, def *defs.SectorDef) *Sector {
	aSector := &Sector{}
	aSector.NearbySectors = make(ConnectionMap)
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

func (s *Sector) AddNearbySector(sector SectorKey, connection SectorConnection) {
	s.NearbySectors[sector] = connection
}

func (s *Sector) GetNearbySectors() ConnectionMap {
	return s.NearbySectors
}

func (s *Sector) GetNearbySector(key SectorKey) SectorConnection {
	return s.NearbySectors[key]
}

func (s *Sector) SpawnBuildings() {
	centerBuilding := NewBuilding(s.Template.CenterBuilding, theWorld.Grid.Tile(s.Center))
	centerBuilding.Set(FactionId, s.Get(FactionId))
	centerBuilding.SpawnAllUnits()
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
		building := NewBuilding(def.Name, theWorld.Grid.Tile(bestLocation))
		building.Set(FactionId, s.Get(FactionId))
		building.SpawnAllUnits()
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
			if theWorld.Grid.Tile(aCoord).HasSector() {
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

func (s *Sector) GetFaction() *Faction {
	return theWorld.GetFaction(FactionKey(s.Get(FactionId)))
}

func (s *Sector) GetEmptyWorkPlace() *Building {
	for _, tileCoord := range s.Tiles {
		tile := theWorld.Grid.Tile(tileCoord)
		building := tile.GetBuilding()
		if building == nil {
			continue
		}
		if building.HasWorkSlot() {
			return building
		}
	}
	return nil
}

func (s *Sector) GetTileNearestTo(tile *Tile) *Tile {
	if len(s.Tiles) == 0 {
		return nil
	}
	tileCoord := tile.GetCoord()
	bestDist := 999999
	bestCoord := tiling.Coord{}
	hasBest := false
	for _, coord := range s.Tiles {
		dist := tileCoord.ChebyshevDist(coord)
		if !hasBest || bestDist > dist {
			hasBest = true
			bestDist = dist
			bestCoord = coord
		}
	}
	if !hasBest {
		return nil
	}
	return theWorld.Grid.Tile(bestCoord)
}

func (s *Sector) SetFaction(faction *Faction) {
	id := int(faction.GetId())
	s.Set(FactionId, id)
	for _, tileCoord := range s.Tiles {
		tile := theWorld.Grid.Tile(tileCoord)
		building := tile.GetBuilding()
		if building == nil {
			continue
		}
		building.Set(FactionId, id)
		for _, unitPointer := range building.Units {
			unit := unitPointer.Get()
			if unit != nil {
				unit.Set(FactionId, id)
			}
		}
	}
}
