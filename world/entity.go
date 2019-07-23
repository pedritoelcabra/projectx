package world

import (
	"github.com/pedritoelcabra/projectx/gfx"
)

type Entity interface {
	DrawSprite(screen *gfx.Screen)
	SetPosition(float64, float64)
	Update(int)
}
