package world

import (
	"github.com/pedritoelcabra/projectx/gfx"
)

type World struct {
	Entities    map[int]Entity
	PlayerUnit  *Player
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
	w.PlayerUnit = NewPlayer()
	w.PlayerUnit.SetPosition(400, 400)
	w.AddEntity(w.PlayerUnit)
	w.initialised = true
}

func (w *World) AddEntity(entity Entity) {
	w.Entities[w.entityCount] = entity
	w.entityCount++
}

func (w *World) Draw(screen *gfx.Screen) {
	if !w.initialised {
		return
	}
	for _, e := range w.Entities {
		e.DrawSprite(screen)
	}
}

func (w *World) Update(tick int) {
	if !w.initialised {
		return
	}
	for _, e := range w.Entities {
		e.Update(tick)
	}
}
