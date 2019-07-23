package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/core"
	"log"
)

func main() {
	handleError(ebiten.Run(update, core.ScreenWidth, core.ScreenHeight, 1, "ProjectX"))
}

func update(screen *ebiten.Image) error {
	return core.G().Update(screen)
}

func handleError(err ...interface{}) {
	if err[0] == nil {
		return
	}
	log.Fatal(err...)
}
