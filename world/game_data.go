package world

import (
	"github.com/pedritoelcabra/projectx/world/grid"
)

type SaveGameData struct {
	Seed   int
	Tick   int
	Player Player
	Grid   grid.Grid
}
