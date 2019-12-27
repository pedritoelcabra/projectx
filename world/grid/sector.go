package grid

import "github.com/pedritoelcabra/projectx/world/coord"

type Sector struct {
	Id     int
	Center coord.Coord
}

func (s *Sector) GetCenter() coord.Coord {
	return s.Center
}

func (s *Sector) GetId() int {
	return s.Id
}
