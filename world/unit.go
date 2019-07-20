package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gfx"
)

type Unit struct {
	sprite gfx.Sprite
	x      float64
	y      float64
}

func (u *Unit) DrawSprite(screen *ebiten.Image) {
	u.sprite.DrawSprite(screen, u.x, u.y)
}

func (u *Unit) SetPosition(x, y float64) {
	u.x = x
	u.y = y
}

func NewUnit() *Unit {
	aUnit := &Unit{}
	aUnit.sprite = gfx.NewLpcSprite(gfx.BodyMaleLight)
	return aUnit
}
