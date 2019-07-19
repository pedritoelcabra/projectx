package world

import (
	"github.com/hajimehoshi/ebiten"
)

type Entity interface {
	Draw(*ebiten.Image)
}
