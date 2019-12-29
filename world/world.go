package world

import (
	"github.com/pedritoelcabra/projectx/core/randomizer"
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/world/tiling"
	"github.com/pedritoelcabra/projectx/world/utils"
)

type EntityMap map[int]Entity

type World struct {
	Entities    EntityMap
	PlayerUnit  *Player
	Grid        *Grid
	entityCount int
	initialised bool
	seed        int
	tick        int
	renderMode  TileRenderMode
}

var theWorld = &World{}

func NewWorld() *World {
	aWorld := &World{}
	aWorld.tick = 0
	theWorld = aWorld
	return aWorld
}

func (w *World) IsInitialized() bool {
	return w.initialised
}

func (w *World) SetSeed(seed int) {
	w.seed = seed
	randomizer.SetSeed(seed)
}

func (w *World) GetSeed() int {
	return w.seed
}

func (w *World) GetTick() int {
	return w.tick
}

func (w *World) Init() {
	tiling.InitTiling()
	utils.Seed(w.seed)
	w.Grid = New()
	w.Entities = make(map[int]Entity)
	w.PlayerUnit = NewPlayer()
	w.PlayerUnit.SetPosition(400, 400)
	w.AddEntity(w.PlayerUnit)
	w.renderMode = RenderModeBasic
	w.initialised = true
}

func LoadFromSave(data SaveGameData) *World {
	w := NewWorld()
	tiling.InitTiling()
	w.SetSeed(data.Seed)
	w.tick = data.Tick
	utils.Seed(w.seed)
	w.Grid = &data.Grid
	w.Entities = make(map[int]Entity)
	w.PlayerUnit = &data.Player
	w.Grid.ChunkGeneration(tiling.NewCoord(tiling.PixelFToTileI(w.PlayerUnit.GetPos())), 0)
	w.PlayerUnit.Unit.InitObjects()
	w.AddEntity(w.PlayerUnit)
	w.renderMode = RenderModeBasic
	w.initialised = true
	return w
}

func (w *World) GetSaveState() SaveGameData {
	state := SaveGameData{}
	state.Seed = w.GetSeed()
	state.Tick = w.GetTick()
	state.Player = *w.PlayerUnit
	state.Grid = *w.Grid
	return state
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
	w.DrawEntities(screen)
}

func (w *World) DrawEntities(screen *gfx.Screen) {
	for _, e := range w.Entities {
		e.DrawSprite(screen)
	}
}

func (w *World) Update() {
	if !w.initialised {
		return
	}
	w.Grid.ChunkGeneration(tiling.NewCoord(tiling.PixelFToTileI(w.PlayerUnit.GetPos())), w.tick)
	for _, e := range w.Entities {
		e.Update(w.tick, w.Grid)
	}
	w.tick++
}

func (w *World) SetRenderMode(mode TileRenderMode) {
	w.renderMode = mode
}
