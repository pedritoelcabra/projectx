package gui

import "github.com/hajimehoshi/ebiten"

type Gui struct {
	game  interface{}
	menus map[string]menu
}

func New(game interface{}) Gui {
	aGui := Gui{}
	aGui.game = game
	aGui.menus["startMenu"] = StartMenu(aGui)
	return aGui
}

func (g Gui) Update(screen *ebiten.Image) {
	for _, menu := range g.menus {
		menu.Update(screen)
	}
}
