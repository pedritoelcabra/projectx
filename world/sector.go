package world

import (
	"github.com/pedritoelcabra/projectx/defs"
	"github.com/pedritoelcabra/projectx/world/container"
	"github.com/pedritoelcabra/projectx/world/tiling"
	"strconv"
)

type SectorKey int

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
	aSector.GrowSector()
	aSector.Init()
	return aSector
}

func (s *Sector) GrowSector() {
	s.Tiles = append(s.Tiles, s.Center)
	s.Tiles = append(s.Tiles, tiling.NewCoord(s.Center.X()-1, s.Center.Y()-1))
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
