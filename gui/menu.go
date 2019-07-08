package gui

import "github.com/hajimehoshi/ebiten"

type menu struct {
	parent     interface{}
	components map[string]interface{}
}

func (m menu) Update(screen *ebiten.Image) {
	return
}

func StartMenu(parent interface{}) menu {
	startMenu := menu{parent, make(map[string]interface{})}
	return startMenu
}
