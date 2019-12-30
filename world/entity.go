package world

import (
	"github.com/pedritoelcabra/projectx/gfx"
)

type EntityKey int

type Entity interface {
	DrawSprite(screen *gfx.Screen)
	SetPosition(float64, float64)
	Update(int, *Grid)
	GetClassName() string
	Init()
}
