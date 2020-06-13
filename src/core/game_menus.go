package core

import (
	"github.com/pedritoelcabra/projectx/src/core/file"
	"github.com/pedritoelcabra/projectx/src/core/world"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/gui"
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
	EntityMenu   string = "entity"
)

func (g *game) InitMenus() {
	g.Gui.AddMenu(StartMenu, g.BuildStartMenu())
	g.Gui.AddMenu(DebugMenu, g.BuildDebugMenu())
	g.Gui.AddMenu(InGameMenu, g.BuildInGameMenu())
	g.Gui.AddMenu(LogMenu, g.BuildLog())
	g.Gui.AddMenu(BuildingMenu, g.BuildBuildings())
	g.Gui.AddMenu(EntityMenu, g.BuildEntityMenu())
}

func (g *game) RebuildInGameMenu() {
	if g.Gui.GetMenu(InGameMenu).IsDisabled() {
		return
	}
	g.Gui.AddMenu(InGameMenu, g.BuildInGameMenu())
	g.Gui.GetMenu(InGameMenu).SetDisabled(false)
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

func (g *game) BuildContextMenu(x, y int) *gui.Menu {
	aMenu := gui.NewMenu(g.Gui)
	if g.isPaused {
		aMenu.SetDisabled(true)
	}

	buttonSize := image.Rect(0, 0, 100, 25)

	mouseCoord := g.MouseTileCoord()
	tile := g.World.Grid.Tile(mouseCoord)
	building := tile.Building.Get()
	if building != nil {
		buildingButton := gui.NewButton(buttonSize, building.GetName())
		buildingButton.OnPressed = func(b *gui.Button) {
			g.ShowEntity(building)
			g.Gui.SetDisabled(ContextMenu, true)
		}
		aMenu.AddButton(buildingButton)
	}

	sector := g.World.GetSector(world.SectorKey(tile.Get(world.SectorId)))

	if sector != nil {
		growButton1 := gui.NewButton(buttonSize, "Grow Sector by 5")
		growButton1.OnPressed = func(b *gui.Button) {
			sector.GrowSectorToSize(5, mouseCoord)
			g.Gui.SetDisabled(ContextMenu, true)
		}
		aMenu.AddButton(growButton1)

		growButton2 := gui.NewButton(buttonSize, "Grow Sector by 8")
		growButton2.OnPressed = func(b *gui.Button) {
			sector.GrowSectorToSize(8, mouseCoord)
			g.Gui.SetDisabled(ContextMenu, true)
		}
		aMenu.AddButton(growButton2)
	}

	unitButton := gui.NewButton(buttonSize, "Add Peasant")
	clickedCoord := g.MousePosCoord()
	unitButton.OnPressed = func(b *gui.Button) {
		world.NewUnit("Peasant", clickedCoord)
		g.Gui.SetDisabled(ContextMenu, true)
	}
	aMenu.AddButton(unitButton)

	unitButton2 := gui.NewButton(buttonSize, "Add Wolf")
	unitButton2.OnPressed = func(b *gui.Button) {
		world.NewUnit("Wolf", clickedCoord)
		g.Gui.SetDisabled(ContextMenu, true)
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

	if g.World != nil && g.World.PlayerUnit.IsInOwnedSector() {
		aButton := gui.NewButton(buttonSize, "Buildings")
		aButton.OnPressed = func(b *gui.Button) {
			g.Gui.ToggleDisabled(BuildingMenu)
		}
		aMenu.AddButton(aButton)
	} else {
		if g.Gui.GetMenu(BuildingMenu) != nil {
			g.Gui.GetMenu(BuildingMenu).SetDisabled(true)
		}
	}

	aMenu.SetDisabled(true)

	return aMenu
}

func InGameMenuHeight() int {
	return 30
}
