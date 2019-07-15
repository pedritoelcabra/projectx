package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gui"
	"strconv"
)

type game struct {
	GUI         *gui.Gui
	tick        int
	framesDrawn int
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
	return nil
}

func (g *game) Update(screen *ebiten.Image) error {
	g.tick++
	g.GUI.Update()

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	g.framesDrawn++

	g.GUI.Draw(screen)
	return nil
}

func (g *game) DebugInfo() string {
	aString := "Tick: " + strconv.Itoa(g.tick)

	return aString
}
