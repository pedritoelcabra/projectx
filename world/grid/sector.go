package grid

import (
	"github.com/pedritoelcabra/projectx/world/container"
	"github.com/pedritoelcabra/projectx/world/coord"
)

type Sector struct {
	Id     int
	Center coord.Coord
	Data   *container.Container
}

func NewSector(location coord.Coord) *Sector {
	aSector := &Sector{}
	aSector.Data = container.NewContainer()
	aSector.Center = location
	return aSector
}

func (s *Sector) GetCenter() coord.Coord {
	return s.Center
}

func (s *Sector) GetId() int {
	return s.Id
}
