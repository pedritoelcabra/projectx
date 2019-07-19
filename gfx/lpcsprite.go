package gfx

import "github.com/hajimehoshi/ebiten"

type LpcSprite struct {
	key spriteKey
}

func NewSprite(key spriteKey) LpcSprite {
	return LpcSprite{key}
}

func (s *LpcSprite) Draw(screen *ebiten.Image) {

}
