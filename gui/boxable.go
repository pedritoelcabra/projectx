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

func (b *button) getBox() image.Rectangle {
	return b.box
}

func (m *menu) getBox() image.Rectangle {
	return m.box
}

func (t *textBox) getBox() image.Rectangle {
	return t.box
}
