package world

import "github.com/hajimehoshi/ebiten"

type Entity interface {
	DrawSprite(*ebiten.Image)
}
