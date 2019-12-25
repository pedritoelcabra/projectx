package world

import (
	"github.com/pedritoelcabra/projectx/core/file"
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/world/grid"
	"github.com/pedritoelcabra/projectx/world/noise"
	"github.com/pedritoelcabra/projectx/world/units"
)

type World struct {
	Entities    map[int]Entity
	PlayerUnit  *units.Player
	Noise       *noise.NoiseGenerator
	Grid        *grid.Grid
	entityCount int
	initialised bool
	seed        int
	tick        int
	renderMode  TileRenderMode
}

func NewWorld() *World {
	aWorld := &World{}
	aWorld.tick = 0
	return aWorld
}

func (w *World) IsInitialized() bool {
	return w.initialised
}

func (w *World) SetSeed(seed int) {
	w.seed = seed
}

func (w *World) GetSeed() int {
	return w.seed
}

func (w *World) GetTick() int {
	return w.tick
}

func (w *World) Init() {
	w.Noise = noise.New(w.seed)
	w.Grid = grid.New()
	w.Grid.SetNoise(w.Noise)
	w.Entities = make(map[int]Entity)
	w.PlayerUnit = units.NewPlayer()
	w.PlayerUnit.SetPosition(400, 400)
	w.AddEntity(w.PlayerUnit)
	w.renderMode = RenderModeBasic
	w.initialised = true
}

func (w *World) LoadFromSave(data file.SaveGameData) {
	w.SetSeed(data.Seed)
	w.tick = data.Tick
	w.Noise = noise.New(w.seed)
	w.Grid = grid.New()
	w.Grid.SetNoise(w.Noise)
	w.Entities = make(map[int]Entity)
	w.PlayerUnit = &data.Player
	w.PlayerUnit.Unit.InitObjects()
	w.AddEntity(w.PlayerUnit)
	w.renderMode = RenderModeBasic
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
	RenderTiles(screen, w, w.renderMode)
	for _, e := range w.Entities {
		e.DrawSprite(screen)
	}
}

func (w *World) Update() {
	if !w.initialised {
		return
	}
	for _, e := range w.Entities {
		e.Update(w.tick)
	}
	w.tick++
}

func (w *World) SetRenderMode(mode TileRenderMode) {
	w.renderMode = mode
}
