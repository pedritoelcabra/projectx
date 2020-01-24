package core

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/file"
	"github.com/pedritoelcabra/projectx/src/core/logger"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/gui"
	"github.com/pedritoelcabra/projectx/src/world"
	"image"
	"image/color"
	"log"
)

const (
	StartMenu    string = "start"
	DebugMenu    string = "debug"
	InGameMenu   string = "game"
	LogMenu      string = "log"
	BuildingMenu string = "building"
	ContextMenu  string = "context"
)

func (g *game) InitMenus() {
	g.Gui.AddMenu(StartMenu, g.BuildStartMenu())
	g.Gui.AddMenu(DebugMenu, g.BuildDebugMenu())
	g.Gui.AddMenu(InGameMenu, g.BuildInGameMenu())
	g.Gui.AddMenu(LogMenu, g.BuildLog())
	g.Gui.AddMenu(BuildingMenu, g.BuildBuildings())
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

	aBox := gui.NewTextBox()
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
	aBox := gui.NewTextBox()
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

func (g *game) BuildBuildings() *gui.Menu {
	buildingMenu := gui.NewMenu(g.Gui)
	buildingMenu.SetHorizontalMenu(true)

	titleSize := 100
	aBox := gui.NewTextBox()
	aBox.SetBox(image.Rect(0, 0, gfx.ScreenWidth, titleSize))
	aBox.SetColor(color.White)
	aBox.SetText("Buildings")
	aBox.SetTextSize(gui.FontSize24)
	aBox.SetHCentered(true)
	buildingMenu.AddTextBox(aBox)

	buttonSize := image.Rect(0, 0, 100, 100)
	for _, def := range defs.BuildingDefs() {
		buildingName := def.Name
		buildingButton := gui.NewButton(buttonSize, buildingName)
		buildingButton.OnPressed = func(b *gui.Button) {
			logger.General(buildingName+" selected", nil)
		}
		buildingMenu.AddButton(buildingButton)
	}

	buildingMenu.SetDisabled(true)
	return buildingMenu
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
	clickedCoord := g.MousePosCoord()
	unitButton.OnPressed = func(b *gui.Button) {
		world.NewUnit("Peasant", clickedCoord)
	}
	aMenu.AddButton(unitButton)

	unitButton2 := gui.NewButton(buttonSize, "Add Wolf")
	unitButton2.OnPressed = func(b *gui.Button) {
		world.NewUnit("Wolf", clickedCoord).Set(world.FactionId, int(g.World.GetFactionByName("Wild Animals").Id))
	}
	aMenu.AddButton(unitButton2)

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

	aButton := gui.NewButton(buttonSize, "Buildings")
	aButton.OnPressed = func(b *gui.Button) {
		g.Gui.ToggleDisabled(BuildingMenu)
	}
	aMenu.AddButton(aButton)

	aMenu.SetDisabled(true)

	return aMenu
}

func InGameMenuHeight() int {
	return 30
}
