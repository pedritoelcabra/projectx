package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gui"
	"strconv"
)

type game struct {
	Gui            *gui.Gui
	Input          *Input
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
