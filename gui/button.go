package gui

import "image"

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

type button struct {
	box     image.Rectangle
	drawBox image.Rectangle
	text    string
}

func (b *button) draw(drawFun drawFunction, box image.Rectangle) {
	b.drawBox.Min.X = box.Min.X + b.box.Min.X
	b.drawBox.Max.X = box.Min.X + b.box.Max.X
	b.drawBox.Min.Y = box.Min.Y + b.box.Min.Y
	b.drawBox.Max.Y = box.Min.Y + b.box.Max.Y
	drawFun(b.drawBox, imageSrcRects[imageTypeButton])
}

func (b *button) getWidth() int {
	return b.box.Max.X - b.box.Min.X
}

func (b *button) getHeight() int {
	return b.box.Max.Y - b.box.Min.Y
}

func (b *button) getBox() image.Rectangle {
	return b.box
}
