package world

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/container"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"github.com/pedritoelcabra/projectx/src/world/utils"
	"log"
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

func (w *World) GetTock() int {
	return w.GetTick() / 60
}

func (w *World) IsTock() bool {
	rest := w.GetTick() % 60
	return rest == 0
}

func FromSeed(seed int) *World {
	w := NewWorld()
	w.Data = container.NewContainer()
	w.Data.Set(Tick, 0)
	w.SetSeed(seed)
	w.Grid = NewGrid()
	w.Entities = NewEntities()
	w.LoadInitialFactions()
	w.PlayerUnit = NewPlayer()
	w.Init()
	w.WorldGeneration()
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
	w.Grid.ChunkGeneration()
	return w
}

func (w *World) LoadInitialFactions() {
	for _, faction := range defs.FactionDefs() {
		NewFaction(faction.Name)
	}
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

func (w *World) WorldGeneration() {
	playerFaction := NewFaction("Player")
	w.PlayerUnit.unit.SetFaction(playerFaction)

	for i := 0; i < 100; i++ {
		if i == 99 {
			log.Fatal("Exceeded 100 iterations in World Gen")
		}
		w.Grid.ChunkGeneration()
		firstSector := w.GetSector(0)
		if firstSector != nil {
			i = 100
			firstSector.SetFaction(playerFaction)
			firstSector.SetName("Idleville")
			w.PlayerUnit.unit.SetF(HomeSector, 0.0)
			w.PlayerUnit.MoveToHomeSector()
		}
	}
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
	drawableBuildings := w.DrawableBuildings()
	for i := 0; i < len(drawableBuildings); i++ {
		drawableBuildings[i].DrawSprite(screen)
	}
	for key, e := range w.DrawableUnits() {
		e.DrawSprite(screen)
		_ = key
	}
}

func (w *World) ClosestUnitWithinRadius(x, y, radius int) *Unit {
	var bestUnit *Unit
	bestDist := 9999
	for _, u := range w.Entities.Units {
		if !u.IsAlive() {
			continue
		}
		ux := utils.AbsInt(int(u.GetX()) - x)
		uy := utils.AbsInt(int(u.GetY()) - y)
		thisDist := utils.MaxInt(uy, ux)
		if thisDist > radius || thisDist > bestDist {
			continue
		}
		bestDist = thisDist
		bestUnit = u
	}
	return bestUnit
}

func (w *World) DrawableBuildings() []*Building {
	var buildings []*Building
	for i := 0; i < len(w.Entities.Buildings); i++ {
		building := w.Entities.Buildings[BuildingKey(i)]
		if building.ShouldDraw() {
			buildings = append(buildings, building)
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
	w.RegisterUnitsWithSectors()
	w.Grid.ChunkGeneration()
	w.PlayerUnit.Update()
	for _, e := range w.Entities.Units {
		e.Update()
	}
	for _, b := range w.Entities.Buildings {
		b.Update()
	}
	w.Grid.Update()
	w.Data.Set(Tick, w.Data.Get(Tick)+1)
}

func (w *World) RegisterUnitsWithSectors() {
	if !w.IsTock() {
		return
	}
	for _, unit := range w.Entities.Units {
		w.Grid.Chunk(w.Grid.ChunkCoord(unit.GetTileCoord())).RegisterUnit(unit)
	}
}

func (w *World) GetScreen() *gfx.Screen {
	return w.screen
}

func (w *World) GetUnits() UnitMap {
	return w.Entities.Units
}
