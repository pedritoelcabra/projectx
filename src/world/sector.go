package world

import (
	"github.com/pedritoelcabra/projectx/defs"
	container2 "github.com/pedritoelcabra/projectx/src/world/container"
	tiling2 "github.com/pedritoelcabra/projectx/src/world/tiling"
	"strconv"
)

type SectorKey int

type Sector struct {
	Id       SectorKey
	Size     int
	Name     string
	Center   tiling2.Coord
	Data     *container2.Container
	Template *defs.SectorDef
	Tiles    []tiling2.Coord
}

func NewSector(location tiling2.Coord, def *defs.SectorDef) *Sector {
	aSector := &Sector{}
	aSector.Template = def
	aSector.Data = container2.NewContainer()
	aSector.Center = location
	aSector.Id = theWorld.AddSector(aSector)
	aSector.Name = aSector.Template.Name + " " + strconv.Itoa(int(aSector.Id))
	aSector.AddTile(aSector.Center)
	aSector.GrowSectorToSize(def.Size)
	aSector.Init()
	return aSector
}

func (s *Sector) GrowSectorToSize(size int) {
	s.Size = size
	options := NewPathOptions()
	options.MinMoveCost = 1.0
	sizeF := float64(size)
	for x := s.Center.X() - size; x <= s.Center.X()+size; x++ {
		for y := s.Center.Y() - size; y <= s.Center.Y()+size; y++ {
			aCoord := tiling2.NewCoord(x, y)
			path := FindPathWithOptions(s.Center, aCoord, options)
			if path.IsValid() && path.GetCost() <= sizeF {
				s.AddTile(aCoord)
			}
		}
	}
}

func (s *Sector) AddTile(tileCoord tiling2.Coord) {
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

func (s *Sector) GetCenter() tiling2.Coord {
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
