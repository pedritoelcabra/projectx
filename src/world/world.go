package world

import (
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/container"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"github.com/pedritoelcabra/projectx/src/world/utils"
	"sort"
)

type World struct {
	Entities    *Entities
	PlayerUnit  *Player
	Grid        *Grid
	Data        *container.Container
	initialised bool
	seed        int
	tick        int
	screen      *gfx.Screen
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
	utils.Seed(w.seed)
}

func (w *World) GetSeed() int {
	return w.seed
}

func (w *World) GetTick() int {
	return w.tick
}

func FromSeed(seed int) *World {
	w := NewWorld()
	w.Data = container.NewContainer()
	w.SetSeed(seed)
	w.Grid = NewGrid()
	w.Entities = NewEntities()

	w.PlayerUnit = NewPlayer()
	w.PlayerUnit.SetPosition(400, 400)
	playerFaction := NewFaction("Player")
	w.PlayerUnit.Set(FactionId, int(playerFaction.GetId()))

	w.Init()
	return w
}

func LoadFromSave(data SaveGameData) *World {
	w := NewWorld()
	w.Data = data.Data
	w.SetSeed(data.Seed)
	w.Grid = &data.Grid
	w.tick = data.Tick
	w.PlayerUnit = &data.Player
	w.Entities = data.WorldEntities
	w.Init()
	w.Grid.ChunkGeneration(tiling.NewCoord(tiling.PixelFToTileI(w.PlayerUnit.GetPos())), 0)
	return w
}

func (w *World) GetSaveState() SaveGameData {
	for _, chunk := range w.Grid.Chunks {
		chunk.PreSave()
	}
	state := SaveGameData{}
	state.Data = w.Data
	state.Seed = w.GetSeed()
	state.Tick = w.GetTick()
	state.Player = *w.PlayerUnit
	state.Grid = *w.Grid
	state.WorldEntities = w.Entities
	return state
}

func (w *World) Init() {
	w.InitEntities()
	tiling.InitTiling()
	w.Data.Set(RenderMode, int(RenderModeBasic))
	w.initialised = true
	InitTileRenderer()
}

func (w *World) InitEntities() {
	w.PlayerUnit.Init()
	for _, unit := range w.Entities.Units {
		unit.Init()
	}
	for _, building := range w.Entities.Buildings {
		building.Init()
	}
	for _, sector := range w.Entities.Sectors {
		sector.Init()
	}
}

func (w *World) Draw(screen *gfx.Screen) {
	if !w.initialised {
		return
	}
	w.screen = screen
	RenderTiles(screen, w, w.GetRenderMode())
	w.DrawEntities(screen)
}

func (w *World) GetRenderMode() TileRenderMode {
	return TileRenderMode(w.Data.Get(RenderMode))
}

func (w *World) SetRenderMode(mode TileRenderMode) {
	w.Data.Set(RenderMode, int(mode))
}

func (w *World) DrawEntities(screen *gfx.Screen) {
	for _, e := range w.Entities.Buildings {
		e.DrawSprite(screen)
	}
	for key, e := range w.DrawableUnits() {
		e.DrawSprite(screen)
		_ = key
	}
}

func (w *World) DrawableUnits() []*Unit {
	var drawableUnits []*Unit
	for _, e := range w.Entities.Units {
		if e.ShouldDraw() {
			drawableUnits = append(drawableUnits, e)
		}
	}
	sort.SliceStable(drawableUnits, func(i, j int) bool {
		return drawableUnits[i].GetY() < drawableUnits[j].GetY()
	})
	return drawableUnits
}

func (w *World) Update() {
	if !w.initialised {
		return
	}
	w.Grid.ChunkGeneration(tiling.NewCoord(tiling.PixelFToTileI(w.PlayerUnit.GetPos())), w.tick)
	w.PlayerUnit.Update(w.tick, w.Grid)
	for _, e := range w.Entities.Units {
		e.Update(w.tick, w.Grid)
	}
	w.tick++
}

func (w *World) GetScreen() *gfx.Screen {
	return w.screen
}
