package core

import (
	"github.com/pedritoelcabra/projectx/src/core/file"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/gui"
	"github.com/pedritoelcabra/projectx/src/world"
	"image"
	"image/color"
	"log"
)

func (g *game) InitMenus() {
	g.Gui.AddMenu("start", g.BuildStartMenu())
	g.Gui.AddMenu("debug", g.BuildDebugMenu())
	g.Gui.AddMenu("game", g.BuildInGameMenu())
	g.Gui.AddMenu("log", g.BuildLog())
}

func (g *game) BuildStartMenu() *gui.Menu {
	aMenu := gui.NewMenu(g.Gui)
	aMenu.SetHCentered(true)
	aMenu.SetTopPadding(50)

	buttonSize := image.Rect(0, 0, 150, 30)

	if g.HasLoadedWorld() {
		continueButton := gui.NewButton(buttonSize, "Resume")
		continueButton.OnPressed = func(b *gui.Button) {
			g.TogglePause()
		}
		aMenu.AddButton(continueButton)
	}

	startButton := gui.NewButton(buttonSize, "New World")
	startButton.OnPressed = func(b *gui.Button) {
		g.InitializeNewWorld()
	}
	aMenu.AddButton(startButton)

	debugButton := gui.NewButton(buttonSize, "Toggle debug")
	debugButton.OnPressed = func(b *gui.Button) {
		g.Gui.ToggleDebug()
	}
	aMenu.AddButton(debugButton)

	if g.HasLoadedWorld() {
		saveButton := gui.NewButton(buttonSize, "Quick Save (F5)")
		saveButton.OnPressed = func(b *gui.Button) {
			g.QuickSave()
		}
		aMenu.AddButton(saveButton)
	}

	if file.SaveGameExists(file.DefaultSaveGameName) {
		loadButton := gui.NewButton(buttonSize, "Quick Load (F9)")
		loadButton.OnPressed = func(b *gui.Button) {
			g.QuickLoad()
		}
		aMenu.AddButton(loadButton)
	}

	stopButton := gui.NewButton(buttonSize, "Exit")
	stopButton.OnPressed = func(b *gui.Button) {
		log.Fatal("Stopped")
	}
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

func (g *game) BuildLog() *gui.Menu {
	logMenu := gui.NewMenu(g.Gui)

	logHeight := 150
	aBox := &gui.TextBox{}
	aBox.SetBox(image.Rect(0, 0, gfx.ScreenWidth, logHeight-InGameMenuHeight()))
	aBox.SetLeftPadding(10)
	aBox.SetTopPadding(gfx.ScreenHeight - logHeight)
	aBox.SetColor(color.White)
	aBox.OnUpdate = func(t *gui.TextBox) {
		t.SetText(g.GetLogText())
	}

	logMenu.AddTextBox(aBox)

	return logMenu
}

func (g *game) BuildContextMenu(x, y int) *gui.Menu {
	aMenu := gui.NewMenu(g.Gui)
	if g.isPaused {
		aMenu.SetDisabled(true)
	}

	buttonSize := image.Rect(0, 0, 100, 25)

	mouseCoord := g.MouseTileCoord()
	tile := g.World.Grid.Tile(mouseCoord)

	sector := g.World.GetSector(world.SectorKey(tile.Get(world.SectorId)))

	if sector != nil {
		growButton1 := gui.NewButton(buttonSize, "Grow Sector by 5")
		growButton1.OnPressed = func(b *gui.Button) {
			sector.GrowSectorToSize(5, mouseCoord)
		}
		aMenu.AddButton(growButton1)

		growButton2 := gui.NewButton(buttonSize, "Grow Sector by 8")
		growButton2.OnPressed = func(b *gui.Button) {
			sector.GrowSectorToSize(8, mouseCoord)
		}
		aMenu.AddButton(growButton2)
	}

	unitButton := gui.NewButton(buttonSize, "Add Peasant")
	unitButton.OnPressed = func(b *gui.Button) {
		world.NewUnit("Peasant", g.MousePosCoord())
	}

	aMenu.AddButton(unitButton)

	aMenu.SetTopPadding(y)
	aMenu.SetLeftPadding(x)

	aMenu.ArrangeContextMenu()

	return aMenu
}

func (g *game) BuildInGameMenu() *gui.Menu {
	aMenu := gui.NewMenu(g.Gui)
	aMenu.SetHorizontalMenu(true)
	aMenu.SetTopPadding(gfx.ScreenHeight - InGameMenuHeight())

	buttonSize := image.Rect(0, 0, 150, InGameMenuHeight())

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

func InGameMenuHeight() int {
	return 30
}
