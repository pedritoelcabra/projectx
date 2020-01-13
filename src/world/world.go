package world

import (
	"encoding/json"
	"github.com/pedritoelcabra/projectx/src/core/randomizer"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"github.com/pedritoelcabra/projectx/src/world/utils"
)

type EntityMap map[EntityKey]Entity
type SectorMap map[SectorKey]*Sector
type FactionMap map[FactionKey]*Faction

type World struct {
	Entities    EntityMap
	Sectors     SectorMap
	Factions    FactionMap
	PlayerUnit  *Player
	Grid        *Grid
	initialised bool
	seed        int
	tick        int
	renderMode  TileRenderMode
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
	w.SetSeed(seed)
	w.Grid = NewGrid()
	w.PlayerUnit = NewPlayer()
	w.PlayerUnit.SetPosition(400, 400)
	w.Entities = make(EntityMap)
	w.Sectors = make(SectorMap)
	w.Init()
	return w
}

func LoadFromSave(data SaveGameData) *World {
	w := NewWorld()
	w.SetSeed(data.Seed)
	w.Grid = &data.Grid
	w.tick = data.Tick
	w.PlayerUnit = &data.Player
	w.Entities = data.Entities
	w.Sectors = data.Sectors
	w.Init()
	w.Grid.ChunkGeneration(tiling.NewCoord(tiling.PixelFToTileI(w.PlayerUnit.GetPos())), 0)
	w.PlayerUnit.Unit.InitObjects()
	return w
}

func (w *World) Init() {
	w.InitEntities()
	tiling.InitTiling()
	w.renderMode = RenderModeBasic
	w.initialised = true
	InitTileRenderer()
}

func (w *World) InitEntities() {
	w.PlayerUnit.Init()
	for _, entity := range w.Entities {
		entity.Init()
	}
	for _, sector := range w.Sectors {
		sector.Init()
	}
}

func (w *World) GetSaveState() SaveGameData {
	for _, chunk := range w.Grid.Chunks {
		chunk.PreSave()
	}
	state := SaveGameData{}
	state.Seed = w.GetSeed()
	state.Tick = w.GetTick()
	state.Player = *w.PlayerUnit
	state.Grid = *w.Grid
	state.Entities = w.Entities
	state.Sectors = w.Sectors
	return state
}

func (w *World) AddEntity(entity Entity) EntityKey {
	key := EntityKey(len(w.Entities))
	w.Entities[key] = entity
	return key
}

func (w *World) GetEntity(key EntityKey) Entity {
	return w.Entities[key]
}

func (w *World) AddSector(sector *Sector) SectorKey {
	key := SectorKey(len(w.Sectors))
	w.Sectors[key] = sector
	return key
}

func (w *World) GetSector(key SectorKey) *Sector {
	if key < 0 {
		return nil
	}
	return w.Sectors[key]
}

func (w *World) AddFaction(sector *Faction) FactionKey {
	key := FactionKey(len(w.Factions))
	w.Factions[key] = sector
	return key
}

func (w *World) GetFaction(key FactionKey) *Faction {
	if key < 0 {
		return nil
	}
	return w.Factions[key]
}

func (w *World) Draw(screen *gfx.Screen) {
	if !w.initialised {
		return
	}
	w.screen = screen
	RenderTiles(screen, w, w.renderMode)
	w.DrawEntities(screen)
}

func (w *World) DrawEntities(screen *gfx.Screen) {
	for _, e := range w.Entities {
		if e.GetClassName() != "Building" {
			e.DrawSprite(screen)
		}
	}
	for _, e := range w.Entities {
		if e.GetClassName() == "Building" {
			e.DrawSprite(screen)
		}
	}
	w.PlayerUnit.DrawSprite(screen)
}

func (w *World) Update() {
	if !w.initialised {
		return
	}
	w.Grid.ChunkGeneration(tiling.NewCoord(tiling.PixelFToTileI(w.PlayerUnit.GetPos())), w.tick)
	w.PlayerUnit.Update(w.tick, w.Grid)
	for _, e := range w.Entities {
		e.Update(w.tick, w.Grid)
	}
	w.tick++
}

func (w *World) SetRenderMode(mode TileRenderMode) {
	w.renderMode = mode
}

func (w *World) GetScreen() *gfx.Screen {
	return w.screen
}

// UnmarshalJSON sets *m to a copy of data.
func (e *EntityMap) UnmarshalJSON(data []byte) error {
	aMap := make(EntityMap)

	var entities map[EntityKey]*json.RawMessage
	if err := json.Unmarshal(data, &entities); err != nil {
		return err
	}

	for index, entity := range entities {
		err, parsedEntity := UnMarshalEntity(entity)
		if err == nil {
			aMap[index] = parsedEntity
		}
	}
	*e = aMap
	return nil
}

func UnMarshalEntity(rawString *json.RawMessage) (error, Entity) {
	entityTypes := make(map[string]Entity)
	entityTypes["Player"] = &Player{}
	entityTypes["Building"] = &Building{}
	entityTypes["Unit"] = &Unit{}

	for className, entity := range entityTypes {
		err := json.Unmarshal(*rawString, entity)
		if err == nil && entity.GetClassName() == className {
			return nil, entity
		}
	}
	err := &json.UnmarshalTypeError{}
	return err, nil
}
