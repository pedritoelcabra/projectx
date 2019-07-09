package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gui"
)

type game struct {
	GUI *gui.Gui
}

const (
	ScreenWidth  = 1200
	ScreenHeight = 900
)

func New() (*game, error) {
	aGame := game{}
	error := aGame.init()
	return &aGame, error
}

func (g *game) init() error {
	g.GUI = gui.New(0, 0, ScreenWidth, ScreenHeight)
	return nil
}

func (g *game) Update(screen *ebiten.Image) error {
	g.GUI.Update()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	g.GUI.Draw(screen)
	return nil
}
