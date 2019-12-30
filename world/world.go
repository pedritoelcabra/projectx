package world

import (
	"encoding/json"
	"github.com/pedritoelcabra/projectx/core/randomizer"
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/world/tiling"
	"github.com/pedritoelcabra/projectx/world/utils"
)

type EntityMap map[EntityKey]Entity

type World struct {
	Entities    EntityMap
	PlayerUnit  *Player
	Grid        *Grid
	entityCount EntityKey
	initialised bool
	seed        int
	tick        int
	renderMode  TileRenderMode
}

var theWorld = &World{}

func NewWorld() *World {
	aWorld := &World{}
	aWorld.tick = 0
	aWorld.entityCount = 1
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
	w.Entities = make(EntityMap)
	w.PlayerUnit = NewPlayer()
	w.PlayerUnit.SetPosition(400, 400)
	w.InitEntities()
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
	w.PlayerUnit = &data.Player
	w.Entities = data.Entities
	w.InitEntities()
	w.Grid.ChunkGeneration(tiling.NewCoord(tiling.PixelFToTileI(w.PlayerUnit.GetPos())), 0)
	w.PlayerUnit.Unit.InitObjects()
	w.renderMode = RenderModeBasic
	w.initialised = true
	return w
}

func (w *World) InitEntities() {
	w.PlayerUnit.Init()
	for _, entity := range w.Entities {
		entity.Init()
	}
}

func (w *World) GetSaveState() SaveGameData {
	state := SaveGameData{}
	state.Seed = w.GetSeed()
	state.Tick = w.GetTick()
	state.Player = *w.PlayerUnit
	state.Grid = *w.Grid
	state.Entities = w.Entities
	return state
}

func (w *World) AddEntity(entity Entity) EntityKey {
	w.Entities[w.entityCount] = entity
	w.entityCount++
	return w.entityCount - 1
}

func (w *World) GetEntity(key EntityKey) Entity {
	return w.Entities[key]
}

func (w *World) Draw(screen *gfx.Screen) {
	if !w.initialised {
		return
	}
	RenderTiles(screen, w, w.renderMode)
	w.DrawEntities(screen)
}

func (w *World) DrawEntities(screen *gfx.Screen) {
	w.PlayerUnit.DrawSprite(screen)
	for _, e := range w.Entities {
		e.DrawSprite(screen)
	}
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
