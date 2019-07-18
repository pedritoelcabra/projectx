package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gui"
	"github.com/pedritoelcabra/projectx/world"
	"strconv"
)

type game struct {
	Gui            *gui.Gui
	Input          *Input
	World          *world.World
	tick           int
	framesDrawn    int
	isPaused       bool
	rightMouseDown bool
	debugMessage   string
}

var projectX *game

const (
	ScreenWidth  = 1200
	ScreenHeight = 900
)

func New() *game {
	aGame := game{}
	aGame.init()
	return &aGame
}

func G() *game {
	if projectX == nil {
		projectX = New()
	}
	return projectX
}

func (g *game) init() error {
	g.World = world.NewWorld()
	g.Input = NewInput()
	g.Gui = gui.New(0, 0, ScreenWidth, ScreenHeight)
	g.InitMenus()
	g.isPaused = true
	return nil
}

func (g *game) Update(screen *ebiten.Image) error {

	g.Gui.Update()
	g.Input.Update()

	if !g.isPaused {
		g.ProcessTick()
	}

	if !ebiten.IsDrawingSkipped() {
		g.framesDrawn++
		g.World.Draw(screen)
		g.Gui.Draw(screen)
	}

	return nil
}

func (g *game) ProcessTick() {
	g.tick++
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

func (g *game) DebugInfo() string {
	aString := "Tick: " + strconv.Itoa(g.tick)
	aString += "\nFrame: " + strconv.Itoa(g.framesDrawn)
	if g.debugMessage != "" {
		aString += "\n" + g.debugMessage
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
	if !g.World.IsInitialized() {
		g.World.Init()
	}
	g.isPaused = false
	g.Gui.SetDisabled("start", true)
	g.Gui.SetDisabled("context", true)
	g.Gui.SetDisabled("game", false)
}

func (g *game) Pause() {
	g.isPaused = true
	g.Gui.SetDisabled("context", true)
	g.Gui.SetDisabled("game", true)
	g.Gui.SetDisabled("start", false)
}
