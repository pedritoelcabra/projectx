package gfx

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	ScreenWidth  = 1200
	ScreenHeight = 800
)

type Screen struct {
	image   *ebiten.Image
	cameraX float64
	cameraY float64
}

func NewScreen() *Screen {
	aScreen := &Screen{}
	return aScreen
}

func (s *Screen) SetScreen(image *ebiten.Image) {
	s.image = image
}

func (s *Screen) SetCameraCoords(x, y float64) {
	s.cameraX = x - (ScreenWidth / 2)
	s.cameraY = y - (ScreenHeight / 2)
}

func (s *Screen) DrawImage(image *ebiten.Image, options *ebiten.DrawImageOptions) {
	options.GeoM.Translate(-s.cameraX, -s.cameraY)
	s.image.DrawImage(image, options)
}
