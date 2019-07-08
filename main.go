package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

var projectX, err = New()

func main() {
	handleError(err)
	handleError(ebiten.Run(update, ScreenWidth, ScreenHeight, 1, "ProjectX"))
}

func update(screen *ebiten.Image) error {
	return projectX.Update(screen)
}

func handleError(err ...interface{}) {
	if err[0] == nil {
		return
	}
	log.Fatal(err)
}
