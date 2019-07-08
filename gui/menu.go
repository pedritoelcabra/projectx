package gui

import "image"

type menu struct {
	parent     interface{}
	components []drawable
}

func (m *menu) Update() {

}

func (m *menu) Draw(draw drawFunction) {
	for _, component := range m.components {
		component.Draw(draw)
	}
}

func newMenu(parent interface{}) *menu {
	return &menu{parent, make([]drawable, 0)}
}

func StartMenu(parent interface{}) *menu {
	startMenu := newMenu(parent)
	startMenu.AddButton("Test", image.Rect(10, 10, 150, 50))
	return startMenu
}

func (m *menu) AddButton(text string, rect image.Rectangle) {
	button1 := &button{
		rect: rect,
		text: text,
	}
	m.components = append(m.components, button1)
}
