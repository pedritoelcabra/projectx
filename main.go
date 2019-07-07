package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

func main() {
	projectX, err := New()
	handleError(err)

	handleError(ebiten.Run(projectX.Update, ScreenWidth, ScreenHeight, 1, "ProjectX"))
}

func handleError(err ...interface{}) {
	if err[0] == nil {
		return
	}
	log.Fatal(err)
}
