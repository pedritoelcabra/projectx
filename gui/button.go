package gui

import (
	"image"
	"image/color"
)

type imageType int

const (
	imageTypeButton imageType = iota
	imageTypeButtonPressed
	imageTypeTextBox
	imageTypeVScollBarBack
	imageTypeVScollBarFront
	imageTypeCheckBox
	imageTypeCheckBoxPressed
	imageTypeCheckBoxMark
)

var imageSrcRects = map[imageType]image.Rectangle{
	imageTypeButton:          image.Rect(0, 0, 16, 16),
	imageTypeButtonPressed:   image.Rect(16, 0, 32, 16),
	imageTypeTextBox:         image.Rect(0, 16, 16, 32),
	imageTypeVScollBarBack:   image.Rect(16, 16, 24, 32),
	imageTypeVScollBarFront:  image.Rect(24, 16, 32, 32),
	imageTypeCheckBox:        image.Rect(0, 32, 16, 48),
	imageTypeCheckBoxPressed: image.Rect(16, 32, 32, 48),
	imageTypeCheckBoxMark:    image.Rect(32, 32, 48, 48),
}

type Button struct {
	box        image.Rectangle
	drawBox    image.Rectangle
	text       string
	textBoxImg *TextBox
	mouseDown  bool
	disabled   bool
	OnPressed  func(b *Button)
}

func NewButton(box image.Rectangle, text string) *Button {
	aButton := &Button{}
	aButton.text = text
	aButton.box = box
	aButton.textBoxImg = &TextBox{}
	aButton.textBoxImg.text = text
	aButton.textBoxImg.box = box
	aButton.textBoxImg.vCenter = true
	aButton.textBoxImg.hCenter = true
	aButton.textBoxImg.fontColor = color.Black
	return aButton
}

func (b *Button) SetDisabled(value bool) {
	b.disabled = value
}

func (b *Button) GetDisabled() bool {
	return b.disabled
}
