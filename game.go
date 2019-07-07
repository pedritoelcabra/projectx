package main

import (
	"github.com/hajimehoshi/ebiten"
)

type game struct {
}

const (
	ScreenWidth  = 420
	ScreenHeight = 600
)

func New() (game, error) {
	return game{}, nil
}

func (g *game) Update(screen *ebiten.Image) error {
	return nil
}
