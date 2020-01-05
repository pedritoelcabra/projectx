package gfx

import (
	"github.com/hajimehoshi/ebiten"
	"image"
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

func (s *Screen) GetCameraCoords() (x, y float64) {
	return s.cameraX, s.cameraY
}

func (s *Screen) GetCameraOffset() (x, y float64) {
	return -s.cameraX, -s.cameraY
}

func (s *Screen) DrawImage(image *ebiten.Image, options *ebiten.DrawImageOptions) {
	options.GeoM.Translate(s.GetCameraOffset())
	s.image.DrawImage(image, options)
}

const (
	a0 = 0x40
	a1 = 0xc0
	a2 = 0xff
)

var pixels = []uint8{
	a0, a1, a1, a0,
	a1, a2, a2, a1,
	a1, a2, a2, a1,
	a0, a1, a1, a0,
}

var brushImage, _ = ebiten.NewImageFromImage(&image.Alpha{
	Pix:    pixels,
	Stride: 4,
	Rect:   image.Rect(0, 0, 4, 4),
}, ebiten.FilterDefault)

func DrawDot(x, y float64, screen *Screen, scale float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x-8.0, y-8.0)
	screen.DrawImage(brushImage, op)
}
