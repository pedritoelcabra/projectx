package core

import (
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/gui"
	"github.com/pedritoelcabra/projectx/world"
	"image"
	"image/color"
	"log"
)

func (g *game) InitMenus() {
	g.Gui.AddMenu("start", g.BuildStartMenu())
	g.Gui.AddMenu("debug", g.BuildDebugMenu())
	g.Gui.AddMenu("game", g.BuildInGameMenu())
}

func (g *game) BuildStartMenu() *gui.Menu {
	aMenu := gui.NewMenu(g.Gui)
	aMenu.SetHCentered(true)
	aMenu.SetTopPadding(50)

	buttonSize := image.Rect(0, 0, 150, 30)

	startButton := gui.NewButton(buttonSize, "Start")
	startButton.OnPressed = func(b *gui.Button) {
		g.TogglePause()
	}

	debugButton := gui.NewButton(buttonSize, "Toggle debug")
	debugButton.OnPressed = func(b *gui.Button) {
		g.Gui.GetMenu("debug").ToggleDisabled()
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
	debugMenu := gui.NewMenu(g.Gui)

	aBox := &gui.TextBox{}
	aBox.SetBox(image.Rect(0, 0, 200, 300))
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
	aMenu := gui.NewMenu(g.Gui)
	if g.isPaused {
		aMenu.SetDisabled(true)
	}

	buttonSize := image.Rect(0, 0, 100, 25)

	aButton := gui.NewButton(buttonSize, "Button 1")
	aButton.OnPressed = func(b *gui.Button) {
		g.debugMessage = "Button 1 pressed"
	}
	bButton := gui.NewButton(buttonSize, "Button 2")
	bButton.OnPressed = func(b *gui.Button) {
		g.debugMessage = "Button 2 pressed"
	}
	cButton := gui.NewButton(buttonSize, "Button 3")
	cButton.OnPressed = func(b *gui.Button) {
		g.debugMessage = "Button 3 pressed"
	}

	aMenu.AddButton(aButton)
	aMenu.AddButton(bButton)
	aMenu.AddButton(cButton)

	aMenu.SetTopPadding(y)
	aMenu.SetLeftPadding(x)

	aMenu.ArrangeContextMenu()

	return aMenu
}

func (g *game) BuildInGameMenu() *gui.Menu {
	aMenu := gui.NewMenu(g.Gui)
	aMenu.SetHorizontalMenu(true)
	menuHeight := 30
	aMenu.SetTopPadding(gfx.ScreenHeight - menuHeight)

	buttonSize := image.Rect(0, 0, 150, menuHeight)

	mainMenuButton := gui.NewButton(buttonSize, "Main Menu")
	mainMenuButton.OnPressed = func(b *gui.Button) {
		g.TogglePause()
	}
	aMenu.AddButton(mainMenuButton)

	aButton := gui.NewButton(buttonSize, "Height Map")
	aButton.OnPressed = func(b *gui.Button) {
		g.World.SetRenderMode(world.RenderModeHeight)
	}
	aMenu.AddButton(aButton)

	bButton := gui.NewButton(buttonSize, "Basic Terrain")
	bButton.OnPressed = func(b *gui.Button) {
		g.World.SetRenderMode(world.RenderModeBasic)
	}
	aMenu.AddButton(bButton)

	aMenu.SetDisabled(true)

	return aMenu
}
