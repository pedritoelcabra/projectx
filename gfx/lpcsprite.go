package gfx

import "github.com/hajimehoshi/ebiten"

type LpcSprite struct {
	key spriteKey
}

func NewLpcSprite(key spriteKey) *LpcSprite {
	return &LpcSprite{key}
}

func (s *LpcSprite) DrawSprite(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(GetSprite(s.key), op)
}
