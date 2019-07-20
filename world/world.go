package world

import (
	"github.com/hajimehoshi/ebiten"
)

type World struct {
	Entities    map[int]Entity
	entityCount int
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
	aUnit := NewUnit()
	aUnit.SetPosition(400, 400)
	w.AddEntity(aUnit)
	w.initialised = true
}

func (w *World) AddEntity(entity Entity) {
	w.Entities[w.entityCount] = entity
	w.entityCount++
}

func (w *World) Draw(screen *ebiten.Image) {
	if !w.initialised {
		return
	}
	for _, e := range w.Entities {
		e.DrawSprite(screen)
	}
}
