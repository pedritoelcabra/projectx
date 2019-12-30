package world

import (
	"github.com/pedritoelcabra/projectx/world/container"
	"github.com/pedritoelcabra/projectx/world/tiling"
)

type SectorKey int

type Sector struct {
	Id     SectorKey
	Center tiling.Coord
	Data   *container.Container
}

func NewSector(location tiling.Coord) *Sector {
	aSector := &Sector{}
	aSector.Data = container.NewContainer()
	aSector.Center = location
	aSector.Id = theWorld.AddSector(aSector)
	aSector.Init()
	return aSector
}

func (s *Sector) GetCenter() tiling.Coord {
	return s.Center
}

func (s *Sector) GetId() SectorKey {
	return s.Id
}

func (s *Sector) Init() {
	theWorld.Grid.Chunk(theWorld.Grid.ChunkCoord(s.Center)).SetSector(s)
}
