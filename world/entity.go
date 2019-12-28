package world

import (
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/world/grid"
)

type Entity interface {
	DrawSprite(screen *gfx.Screen)
	SetPosition(float64, float64)
	Update(int, *grid.Grid)
}
