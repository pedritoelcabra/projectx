package gui

import (
	"image"
)

type boxable interface {
	getBox() image.Rectangle
}

func (g *Gui) getBox() image.Rectangle {
	return g.box
}

func (b *Button) getBox() image.Rectangle {
	return b.box
}

func (m *Menu) getBox() image.Rectangle {
	return m.box
}

func (t *TextBox) getBox() image.Rectangle {
	return t.box
}
