package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gfx"
)

type Unit struct {
	sprite gfx.Sprite
}

func (u *Unit) DrawSprite(screen *ebiten.Image) {
	u.sprite.DrawSprite(screen)
}

func NewUnit() *Unit {
	aUnit := &Unit{}
	aUnit.sprite = gfx.NewLpcSprite(gfx.BodyMaleLight)
	return aUnit
}
