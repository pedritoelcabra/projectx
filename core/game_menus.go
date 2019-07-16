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
	g.GUI.AddMenu("game", g.BuildInGameMenu())
}

func (g *game) BuildStartMenu() *gui.Menu {
	aMenu := gui.NewMenu(g.GUI)
	aMenu.SetHCentered(true)
	aMenu.SetTopPadding(50)

	buttonSize := image.Rect(0, 0, 150, 30)

	startButton := gui.NewButton(buttonSize, "Start")
	startButton.OnPressed = func(b *gui.Button) {
		g.isPaused = false
		g.GUI.DisableAllMenus()
		g.GUI.GetMenu("game").SetDisabled(false)
	}

	debugButton := gui.NewButton(buttonSize, "Toggle debug")
	debugButton.OnPressed = func(b *gui.Button) {
		g.GUI.GetMenu("debug").ToggleDisabled()
	}

	stopButton := gui.NewButton(buttonSize, "Exit")
	stopButton.OnPressed = func(b *gui.Button) {
		log.Fatal("Stopped")
	}

	aMenu.AddButton(startButton)
	aMenu.AddButton(debugButton)
	aMenu.AddButton(stopButton)

	return aMenu
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

func (g *game) BuildContextMenu(x, y int) *gui.Menu {
	aMenu := gui.NewMenu(g.GUI)

	buttonSize := image.Rect(0, 0, 100, 25)

	aButton := gui.NewButton(buttonSize, "Button 1")
	bButton := gui.NewButton(buttonSize, "Button 1")
	cButton := gui.NewButton(buttonSize, "Button 1")

	aMenu.AddButton(aButton)
	aMenu.AddButton(bButton)
	aMenu.AddButton(cButton)

	aMenu.SetTopPadding(y)
	aMenu.SetLeftPadding(x)

	aMenu.ArrangeContextMenu()

	return aMenu
}

func (g *game) BuildInGameMenu() *gui.Menu {
	aMenu := gui.NewMenu(g.GUI)
	aMenu.SetHorizontalMenu(true)
	menuHeight := 30
	aMenu.SetTopPadding(ScreenHeight - menuHeight)

	buttonSize := image.Rect(0, 0, 150, menuHeight)

	mainMenuButton := gui.NewButton(buttonSize, "Main Menu")
	mainMenuButton.OnPressed = func(b *gui.Button) {
		g.isPaused = true
		g.GUI.DisableAllMenus()
		g.GUI.GetMenu("start").SetDisabled(false)
	}
	aMenu.AddButton(mainMenuButton)

	aButton := gui.NewButton(buttonSize, "Placeholder 1")
	aMenu.AddButton(aButton)

	bButton := gui.NewButton(buttonSize, "Placeholder 2")
	aMenu.AddButton(bButton)

	aMenu.SetDisabled(true)

	return aMenu
}
