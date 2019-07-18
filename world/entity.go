package world

import "github.com/hajimehoshi/ebiten"

type Entity interface {
	Draw(*ebiten.Image)
}

func (s *LpcSprite) Draw(screen *ebiten.Image) {

}
