package core

import (
	"github.com/pedritoelcabra/projectx/gui"
	"image"
	"image/color"
	"log"
)

func (g *game) InitMenus() {
	g.GUI.AddMenu("start", g.BuildStartMenu())
	g.GUI.AddMenu("debug", g.BuildDebugMenu())
}

func (g *game) BuildStartMenu() *gui.Menu {
	startMenu := gui.NewMenu(g.GUI)
	startMenu.SetHCentered(true)
	startMenu.SetTopPadding(50)

	buttonSize := image.Rect(0, 0, 150, 30)

	stopButton := gui.NewButton(buttonSize, "Stop!")
	stopButton.OnPressed = func(b *gui.Button) {
		log.Fatal("Stopped")
	}

	startMenu.AddButton(stopButton)
	startMenu.SetDisabled(true)

	return startMenu
}

func (g *game) BuildDebugMenu() *gui.Menu {
	debugMenu := gui.NewMenu(g.GUI)

	aBox := &gui.TextBox{}
	aBox.SetBox(image.Rect(0, 0, 200, 120))
	aBox.SetLeftPadding(10)
	aBox.SetTopPadding(10)
	aBox.SetColor(color.White)
	aBox.OnUpdate = func(t *gui.TextBox) {
		t.SetText(g.DebugInfo())
	}

	debugMenu.AddTextBox(aBox)

	return debugMenu
}
