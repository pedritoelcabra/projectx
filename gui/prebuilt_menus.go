package gui

import "image"

func StartMenu(parent interface{}) *menu {
	startMenu := newMenu(parent)
	startMenu.topPadding = 50
	startMenu.AddButton("Test", image.Rect(0, 0, 150, 30))
	startMenu.AddButton("Test2", image.Rect(0, 0, 150, 30))
	return startMenu
}
