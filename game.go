package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gui"
)

type game struct {
	GUI gui.Gui
}

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func New() (game, error) {
	aGame := game{}
	if error := aGame.init(); error != nil {
		return aGame, error
	}
	return aGame, nil
}

func (g game) init() error {
	g.GUI = gui.New(&g)
	return nil
}

func (g *game) Update(screen *ebiten.Image) error {
	g.GUI.Update(screen)
	return nil
}
