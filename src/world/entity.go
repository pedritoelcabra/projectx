package world

import (
	"github.com/pedritoelcabra/projectx/src/gfx"
)

type EntityKey int
type EntityMap map[EntityKey]Entity

type Entity interface {
	DrawSprite(screen *gfx.Screen)
	SetPosition(float64, float64)
	Update(int, *Grid)
	GetClassName() string
	Init()
}
