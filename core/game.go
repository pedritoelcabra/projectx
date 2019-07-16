package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gui"
	"image/color"
	"strconv"
)

type game struct {
	GUI            *gui.Gui
	tick           int
	framesDrawn    int
	isPaused       bool
	rightMouseDown bool
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
	g.GUI = gui.New(0, 0, ScreenWidth, ScreenHeight)
	g.InitMenus()
	g.isPaused = true
	return nil
}

func (g *game) Update(screen *ebiten.Image) error {

	if !g.isPaused {
		g.ProcessTick()
	}

	screen.Fill(color.Black)

	g.GUI.Update()
	g.openContextMenu()

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	g.framesDrawn++

	g.GUI.Draw(screen)
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
			g.GUI.AddMenu("context", g.BuildContextMenu(ebiten.CursorPosition()))
		}
		g.rightMouseDown = false
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.rightMouseDown = false
		contextMenu := g.GUI.GetMenu("context")
		if contextMenu != nil {
			contextMenu.SetDisabled(true)
		}
	}
}

func (g *game) DebugInfo() string {
	aString := "Tick: " + strconv.Itoa(g.tick)

	return aString
}
