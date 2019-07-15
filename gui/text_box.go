package gui

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	"image/color"
)

type TextBox struct {
	box         image.Rectangle
	drawBox     image.Rectangle
	text        string
	contentBuf  *ebiten.Image
	vCenter     bool
	hCenter     bool
	leftPadding int
	topPadding  int
	fontColor   color.Gray16
	OnUpdate    func(t *TextBox)
}

func (t *TextBox) SetBox(box image.Rectangle) {
	t.box = box
}

func (t *TextBox) SetLeftPadding(value int) {
	t.leftPadding = value
}

func (t *TextBox) SetTopPadding(value int) {
	t.topPadding = value
}

func (t *TextBox) SetColor(color color.Gray16) {
	t.fontColor = color
}

func (t *TextBox) SetText(text string) {
	t.text = text
}
