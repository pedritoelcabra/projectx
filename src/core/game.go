package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/file"
	"github.com/pedritoelcabra/projectx/src/core/logger"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/core/world"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/gui"
	"strconv"
)

type game struct {
	Gui              *gui.Gui
	Input            *Input
	World            *world.World
	Graphics         *gfx.Graphics
	Screen           *gfx.Screen
	PlacementManager *PlacementManager
	framesDrawn      int
	isPaused         bool
	rightMouseDown   bool
	debugMessage     string
}

var ProjectX *game

func New() *game {
	aGame := game{}
	aGame.init()
	return &aGame
}

func G() *game {
	if ProjectX == nil {
		ProjectX = New()
	}
	return ProjectX
}

func (g *game) init() {
	logger.InitLogger()
	g.PlacementManager = NewPlacementManager()
	g.Screen = gfx.NewScreen()
	g.Graphics = gfx.NewGraphics()
	defs.InitDefs()
	g.InitInput()
	g.Gui = gui.New(0, 0, gfx.ScreenWidth, gfx.ScreenHeight)
	g.InitMenus()
	g.isPaused = true
	logger.General("Initialised Game", nil)
}

func (g *game) Update(screen *ebiten.Image) error {

	g.Screen.SetScreen(screen)
	g.Gui.Update()
	g.Input.Update()

	if !g.isPaused {
		g.World.Update()
		g.UpdateMenus()
		g.Screen.SetCameraCoords(g.World.PlayerUnit.GetPos())
	}

	if !ebiten.IsDrawingSkipped() {
		if !g.isPaused {
			g.World.Draw(g.Screen)
		}
		g.PlacementManager.Draw(g.Screen)
		g.Gui.Draw(screen)
		g.framesDrawn++
	}

	return nil
}

func (g *game) UpdateMenus() {
	if g.World.PlayerUnit.GetSectorUpdate() {
		g.RebuildInGameMenu()
	}
	displayEntity := g.World.DisplayEntity()
	if displayEntity != nil && g.World.ShouldDisplayEntity() {
		g.ShowEntity(displayEntity)
		g.World.MarkEntityDisplayed()
	}
}

func (g *game) openContextMenu() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.rightMouseDown = true
	} else {
		if g.rightMouseDown {
			g.Gui.AddMenu("context", g.BuildContextMenu(ebiten.CursorPosition()))
		}
		g.rightMouseDown = false
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.rightMouseDown = false
		contextMenu := g.Gui.GetMenu("context")
		if contextMenu != nil {
			contextMenu.SetDisabled(true)
		}
	}
}

func (g *game) GetLogText() string {
	aString := ""
	log := logger.Get(logger.GeneralLog, 6)
	for i := len(log) - 1; i >= 0; i-- {
		if aString != "" {
			aString += "\n"
		}
		aString += log[i].Message()
	}

	return aString
}

func (g *game) TogglePause() {
	if g.isPaused {
		g.UnPause()
		return
	}
	g.Pause()
}

func (g *game) UnPause() {
	g.isPaused = false
	g.Gui.SetDisabled(StartMenu, true)
	g.Gui.SetDisabled(ContextMenu, true)
	g.Gui.SetDisabled(InGameMenu, false)
}

func (g *game) HasLoadedWorld() bool {
	return g.World != nil && g.World.IsInitialized()
}

func (g *game) InitializeNewWorld() {
	g.World = world.FromSeed(randomizer.NewSeed())
	g.InitMenus()
	g.UnPause()
	logger.General("Created a New World with seed "+strconv.Itoa(g.World.GetSeed()), nil)
}

func (g *game) Pause() {
	g.isPaused = true
	g.Gui.SetDisabled(ContextMenu, true)
	g.Gui.SetDisabled(InGameMenu, true)
	g.Gui.SetDisabled(BuildingMenu, true)
	g.Gui.SetDisabled(StartMenu, false)
}

func (g *game) UpdatePlayerMovement(dir world.PlayerDirection, value bool) {
	if g.World == nil || g.World.PlayerUnit == nil {
		return
	}
	g.World.PlayerUnit.SetMovement(dir, value)
}

func (g *game) HandleAttackClick() {
	if g.World == nil || !g.World.IsInitialized() || g.isPaused {
		return
	}
	g.World.PlayerUnit.SetAttackPoint(g.MousePosCoord().XY())
}

func (g *game) QuickSave() {
	file.SaveToFile(g.World.GetSaveState(), file.DefaultSaveGameName)
	g.InitMenus()
	g.UnPause()
	logger.General("Quick Saved", nil)
}

func (g *game) QuickLoad() {
	dataStructure := file.LoadFromFile(file.DefaultSaveGameName)
	g.World = world.LoadFromSave(dataStructure)
	g.InitMenus()
	g.UnPause()
	logger.General("Quick Loaded", nil)
}
