package world

import "github.com/hajimehoshi/ebiten"

type World struct {
	Entities map[int]Entity
}

func NewWorld() *World {
	aWorld := &World{}
	aWorld.Init()
	return aWorld
}

func (w *World) Init() {
	w.Entities = make(map[int]Entity)
}

func (w *World) Draw(screen *ebiten.Image) {
	for _, e := range w.Entities {
		e.Draw(screen)
	}
}
