package gui

import "image"

type menu struct {
	parent      interface{}
	components  []drawable
	box         image.Rectangle
	hCentered   bool
	topPadding  int
	leftPadding int
}

func (m *menu) Update() {

}

func (m *menu) getBox() image.Rectangle {
	return m.box
}

func (m *menu) draw(drawFun drawFunction, box image.Rectangle) {
	drawSpace := box
	drawSpace.Min.X += m.centeredOffset(box)
	for _, component := range m.components {
		component.draw(drawFun, drawSpace)
	}
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

func (m *menu) getWidth() int {
	maxWidth := 0
	for _, component := range m.components {
		if component.getWidth() > maxWidth {
			maxWidth = component.getWidth()
		}
	}
	return maxWidth
}

func newMenu(parent interface{}) *menu {
	aMenu := &menu{}
	aMenu.parent = parent
	aMenu.components = make([]drawable, 0)
	aMenu.hCentered = true
	return aMenu
}

func StartMenu(parent interface{}) *menu {
	startMenu := newMenu(parent)
	startMenu.AddButton("Test", image.Rect(0, 0, 150, 50))
	return startMenu
}

func (m *menu) AddButton(text string, rect image.Rectangle) {
	button1 := &button{
		box:  rect,
		text: text,
	}
	m.components = append(m.components, button1)
}
