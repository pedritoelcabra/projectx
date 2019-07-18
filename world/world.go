package world

import "github.com/hajimehoshi/ebiten"

type World struct {
	Entities    map[int]Entity
	initialised bool
}

func NewWorld() *World {
	aWorld := &World{}
	return aWorld
}

func (w *World) IsInitialized() bool {
	return w.initialised
}

func (w *World) Init() {
	w.Entities = make(map[int]Entity)
	w.initialised = true
}

func (w *World) Draw(screen *ebiten.Image) {
	if !w.initialised {
		return
	}
	for _, e := range w.Entities {
		e.Draw(screen)
	}
}
