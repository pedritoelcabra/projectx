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
	Entities   *Entities
	PlayerUnit *Player
	Grid       *Grid
	Data       *container.Container
	screen     *gfx.Screen
}

var theWorld = &World{}

func NewWorld() *World {
	aWorld := &World{}
	theWorld = aWorld
	return aWorld
}

func (w *World) IsInitialized() bool {
	return w.Data.Get(Initialized) == 1
}

func (w *World) SetSeed(seed int) {
	w.Data.Set(Seed, seed)
	randomizer.SetSeed(seed)
	utils.Seed(seed)
}

func (w *World) GetSeed() int {
	return w.Data.Get(Seed)
}

func (w *World) GetTick() int {
	return w.Data.Get(Tick)
}

func FromSeed(seed int) *World {
	w := NewWorld()
	w.Data = container.NewContainer()
	w.Data.Set(Tick, 0)
	w.SetSeed(seed)
	w.Grid = NewGrid()
	w.Entities = NewEntities()

	w.PlayerUnit = NewPlayer()
	w.PlayerUnit.SetPosition(400, 400)
	playerFaction := NewFaction("Player")
	w.PlayerUnit.unit.SetFaction(playerFaction)

	w.Init()
	return w
}

func LoadFromSave(data SaveGameData) *World {
	w := NewWorld()
	w.Data = data.Data
	w.SetSeed(data.Seed)
	w.Grid = &data.Grid
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
	w.Data.Set(Initialized, 1)
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
	if !w.IsInitialized() {
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
	for _, e := range w.DrawableBuildings() {
		e.DrawSprite(screen)
	}
	for key, e := range w.DrawableUnits() {
		e.DrawSprite(screen)
		_ = key
	}
}

func (w *World) DrawableBuildings() []*Building {
	var buildings []*Building
	for _, e := range w.Entities.Buildings {
		if e.ShouldDraw() {
			buildings = append(buildings, e)
		}
	}
	w.Data.Set(CurrentDrawnBuildings, len(buildings))
	return buildings
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
	w.Data.Set(CurrentDrawnUnits, len(drawableUnits))
	return drawableUnits
}

func (w *World) Update() {
	if !w.IsInitialized() {
		return
	}
	w.Grid.ChunkGeneration(tiling.NewCoord(tiling.PixelFToTileI(w.PlayerUnit.GetPos())), w.GetTick())
	w.PlayerUnit.Update(w.GetTick(), w.Grid)
	for _, e := range w.Entities.Units {
		e.Update(w.GetTick(), w.Grid)
	}
	w.Data.Set(Tick, w.Data.Get(Tick)+1)
}

func (w *World) GetScreen() *gfx.Screen {
	return w.screen
}

func (w *World) GetUnits() UnitMap {
	return w.Entities.Units
}
