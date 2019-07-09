package gui

import "image"

type menu struct {
	parent            interface{}
	components        []drawable
	box               image.Rectangle
	hCentered         bool
	topPadding        int
	leftPadding       int
	horizontalSpacing int
}

func (m *menu) Update() {

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

func newMenu(parent interface{}) *menu {
	aMenu := &menu{}
	aMenu.parent = parent
	aMenu.components = make([]drawable, 0)
	aMenu.hCentered = true
	return aMenu
}

func (m *menu) AddButton(text string, box image.Rectangle) {
	m.components = append(m.components, NewButton(box, text))
}
