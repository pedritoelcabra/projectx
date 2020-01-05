package gui

import "image"

type Menu struct {
	gui            *Gui
	components     []drawable
	box            image.Rectangle
	hCentered      bool
	topPadding     int
	leftPadding    int
	horizontalMenu bool
	disabled       bool
}

func (m *Menu) SetHCentered(centered bool) {
	m.hCentered = centered
}

func (m *Menu) SetTopPadding(padding int) {
	m.topPadding = padding
}

func (m *Menu) SetDisabled(value bool) {
	m.disabled = value
}

func (m *Menu) ToggleDisabled() {
	m.disabled = !m.disabled
}

func (m *Menu) SetHorizontalMenu(value bool) {
	m.horizontalMenu = value
}

func (m *Menu) SetLeftPadding(padding int) {
	m.leftPadding = padding
}

func (m *Menu) centeredOffset(box image.Rectangle) int {
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

func NewMenu(gui *Gui) *Menu {
	aMenu := &Menu{}
	aMenu.gui = gui
	aMenu.components = make([]drawable, 0)
	return aMenu
}

func (m *Menu) AddButton(aButton *Button) {
	m.components = append(m.components, aButton)
}

func (m *Menu) AddTextBox(aBox *TextBox) {
	m.components = append(m.components, aBox)
}

func (m *Menu) ArrangeContextMenu() {
	maxW := m.gui.box.Max.X
	maxH := m.gui.box.Max.Y
	currentX := m.leftPadding
	currentY := m.topPadding
	currentW := m.getWidth()
	currentH := m.getHeight()
	if currentW+currentX > maxW {
		m.leftPadding -= currentW
	}
	if currentH+currentY > maxH {
		m.topPadding -= currentH
	}
	return
}
