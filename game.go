package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gui"
)

type game struct {
	GUI *gui.Gui
}

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func New() (*game, error) {
	aGame := game{}
	error := aGame.init()
	return &aGame, error
}

func (g *game) init() error {
	g.GUI = gui.New(g)
	return nil
}

func (g *game) Update(screen *ebiten.Image) error {
	g.GUI.Update()
	g.GUI.Draw(screen)
	return nil
}
