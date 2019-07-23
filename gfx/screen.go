package gfx

import "github.com/hajimehoshi/ebiten"

type Screen struct {
	image *ebiten.Image
}

func NewScreen() *Screen {
	aScreen := &Screen{}
	return aScreen
}

func (s *Screen) SetScreen(image *ebiten.Image) {
	s.image = image
}

func (s *Screen) DrawImage(image *ebiten.Image, options *ebiten.DrawImageOptions) {
	s.image.DrawImage(image, options)
}
