package gui

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	"image/color"
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
	fontColor   color.Gray16
	onUpdate    func(t *textBox)
}
