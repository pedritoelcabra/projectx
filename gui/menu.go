package gui

import "image"

type menu struct {
	gui               *Gui
	components        []drawable
	box               image.Rectangle
	hCentered         bool
	topPadding        int
	leftPadding       int
	horizontalSpacing int
}

func (m *menu) centeredOffset(box image.Rectangle) int {
	if !m.hCentered {
		return 0
	}
	maxWidth := m.getWidth()
	maxSpace := box.Max.X - box.Min.X
	if maxWidth >= maxSpace {
		return 0
	}
	return (maxSpace - maxWidth) / 2
}

func newMenu(gui *Gui) *menu {
	aMenu := &menu{}
	aMenu.gui = gui
	aMenu.components = make([]drawable, 0)
	return aMenu
}

func (m *menu) addButton(aButton *button) {
	m.components = append(m.components, aButton)
}

func (m *menu) addTextBox(aBox *textBox) {
	m.components = append(m.components, aBox)
}
