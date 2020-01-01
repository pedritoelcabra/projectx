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
	Name     string
	Center   tiling.Coord
	Data     *container.Container
	Template *defs.SectorDef
}

func NewSector(location tiling.Coord, def *defs.SectorDef) *Sector {
	aSector := &Sector{}
	aSector.Template = def
	aSector.Data = container.NewContainer()
	aSector.Center = location
	aSector.Id = theWorld.AddSector(aSector)
	aSector.Name = aSector.Template.Name + " " + strconv.Itoa(int(aSector.Id))
	theWorld.Grid.Tile(location).Set(SectorId, int(aSector.Id))
	aSector.Init()
	return aSector
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
}
