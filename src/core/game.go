package core

import (
	"github.com/hajimehoshi/ebiten"
	defs2 "github.com/pedritoelcabra/projectx/src/core/defs"
	file2 "github.com/pedritoelcabra/projectx/src/core/file"
	logger2 "github.com/pedritoelcabra/projectx/src/core/logger"
	randomizer2 "github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/gui"
	"github.com/pedritoelcabra/projectx/src/world"
	"strconv"
)

type game struct {
	Gui            *gui.Gui
	Input          *Input
	World          *world.World
	Graphics       *gfx.Graphics
	Screen         *gfx.Screen
	framesDrawn    int
	isPaused       bool
	rightMouseDown bool
	debugMessage   string
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
	logger2.InitLogger()
	defs2.InitDefs()
	g.Screen = gfx.NewScreen()
	g.Graphics = gfx.NewGraphics()
	g.InitInput()
	g.Gui = gui.New(0, 0, gfx.ScreenWidth, gfx.ScreenHeight)
	g.InitMenus()
	g.isPaused = true
	logger2.General("Initialised Game", nil)
}

func (g *game) Update(screen *ebiten.Image) error {

	g.Screen.SetScreen(screen)
	g.Gui.Update()
	g.Input.Update()

	if !g.isPaused {
		g.World.Update()
		g.Screen.SetCameraCoords(g.World.PlayerUnit.GetPos())
	}

	if !ebiten.IsDrawingSkipped() {
		if !g.isPaused {
			g.World.Draw(g.Screen)
		}
		g.Gui.Draw(screen)
		g.framesDrawn++
	}

	return nil
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
	log := logger2.Get(logger2.GeneralLog, 6)
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
	g.Gui.SetDisabled("start", true)
	g.Gui.SetDisabled("context", true)
	g.Gui.SetDisabled("game", false)
}

func (g *game) HasLoadedWorld() bool {
	return g.World != nil && g.World.IsInitialized()
}

func (g *game) InitializeNewWorld() {
	g.World = world.NewWorld()
	g.World.SetSeed(randomizer2.NewSeed())
	g.World.Init()
	g.InitMenus()
	g.UnPause()
	logger2.General("Created a New World with seed "+strconv.Itoa(g.World.GetSeed()), nil)
}

func (g *game) Pause() {
	g.isPaused = true
	g.Gui.SetDisabled("context", true)
	g.Gui.SetDisabled("game", true)
	g.Gui.SetDisabled("start", false)
}

func (g *game) UpdatePlayerMovement(dir world.PlayerDirection, value bool) {
	if g.World == nil || g.World.PlayerUnit == nil {
		return
	}
	g.World.PlayerUnit.SetMovement(dir, value)
}

func (g *game) QuickSave() {
	file2.SaveToFile(g.World.GetSaveState(), file2.DefaultSaveGameName)
	g.InitMenus()
	g.UnPause()
	logger2.General("Quick Saved", nil)
}

func (g *game) QuickLoad() {
	dataStructure := file2.LoadFromFile(file2.DefaultSaveGameName)
	g.World = world.LoadFromSave(dataStructure)
	g.InitMenus()
	g.UnPause()
	logger2.General("Quick Loaded", nil)
}
