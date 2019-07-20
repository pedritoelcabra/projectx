package gfx

import (
	"github.com/hajimehoshi/ebiten"
)

type Sprite interface {
	DrawSprite(*ebiten.Image, float64, float64)
}
