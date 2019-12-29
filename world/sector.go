package world

import (
	"github.com/pedritoelcabra/projectx/world/container"
	"github.com/pedritoelcabra/projectx/world/tiling"
)

type Sector struct {
	Id     int
	Center tiling.Coord
	Data   *container.Container
}

func NewSector(location tiling.Coord) *Sector {
	aSector := &Sector{}
	aSector.Data = container.NewContainer()
	aSector.Center = location
	return aSector
}

func (s *Sector) GetCenter() tiling.Coord {
	return s.Center
}

func (s *Sector) GetId() int {
	return s.Id
}
