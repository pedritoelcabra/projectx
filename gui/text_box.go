package gui

import (
	"github.com/hajimehoshi/ebiten"
	"image"
)

type textBox struct {
	box         image.Rectangle
	drawBox     image.Rectangle
	text        string
	contentBuf  *ebiten.Image
	vCenter     bool
	hCenter     bool
	leftPadding int
	topPadding  int
}
