package world

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
	"strconv"
)

type BuildingKey int
type BuildingMap map[BuildingKey]*Building

type Building struct {
	Id         BuildingKey
	Sprite     gfx.Sprite `json:"-"`
	SpriteKey  gfx.SpriteKey
	Name       string
	Location   tiling.Coord
	X          float64
	Y          float64
	Template   *defs.BuildingDef
	Units      []UnitPointer
	Attributes *Attributes
}

func NewBuilding(name string, location *Tile) *Building {
	buildingDefs := defs.BuildingDefs()
	aBuilding := &Building{}
	aBuilding.Template = buildingDefs[name]
	aBuilding.Attributes = NewEmptyAttributes()
	aBuilding.Set(ConstructionProgress, aBuilding.Template.ConstructionWork)
	aBuilding.Set(UnitSpawnProgress, 0)
	aBuilding.SpriteKey = gfx.GetSpriteKey(aBuilding.Template.Graphic)
	aBuilding.Name = name
	aBuilding.Init()
	aBuilding.Location = location.GetCoord()
	aBuilding.X = location.GetF(RenderX)
	aBuilding.Y = location.GetF(RenderY)
	aBuilding.Id = theWorld.AddBuilding(aBuilding)
	location.SetBuilding(aBuilding)
	return aBuilding
}

func (b *Building) Init() {
	b.Sprite = gfx.NewHexSprite(b.SpriteKey)
	tile := theWorld.Grid.Tile(b.Location)
	if tile != nil {
		tile.SetBuilding(b)
	}
}

func (b *Building) SetConstructionProgress(value int) {
	if value < 0 {
		value = 0
	}
	if value > b.Template.ConstructionWork {
		value = b.Template.ConstructionWork
	}
	b.Set(ConstructionProgress, value)
}

func (b *Building) GetConstructionProgress() int {
	return b.Get(ConstructionProgress)
}

func (b *Building) AddConstructionProgress(value int) {
	b.SetConstructionProgress(value + b.GetConstructionProgress())
}

func (b *Building) ConstructionIsComplete() bool {
	return b.Get(ConstructionProgress) >= b.Template.ConstructionWork
}

func (b *Building) CompleteConstruction() {
	b.Set(ConstructionProgress, b.Template.ConstructionWork)
	b.SpriteKey = gfx.GetSpriteKey(b.Template.Graphic)
	b.Init()
}

func (b *Building) StartConstruction() {
	b.Set(ConstructionProgress, 0)
	b.SpriteKey = gfx.GetSpriteKey(b.Template.ConstructionGraphic)
	b.Init()
}

func (b *Building) ShouldDraw() bool {
	return EntityShouldDraw(b.GetX(), b.GetY())
}

func (b *Building) GetX() float64 {
	return b.X
}

func (b *Building) GetY() float64 {
	return b.Y
}

func (b *Building) DrawSprite(screen *gfx.Screen) {
	b.Sprite.DrawSprite(screen, b.X, b.Y)
}

func (b *Building) SetPosition(x, y float64) {

}

func (b *Building) Update() {
	b.UpdateUnitSpawn()
}

func (b *Building) UpdateUnitSpawn() {
	if !b.ConstructionIsComplete() {
		return
	}
	if !theWorld.IsTock() {
		return
	}
	if len(b.Units) >= b.Template.UnitLimit {
		return
	}
	b.AddUnitSpawnProgress(1)
	if b.UnitSpawnIsComplete() {
		b.SpawnUnit()
	}
}

func (b *Building) SpawnUnit() {
	b.Set(UnitSpawnProgress, 0)
	unit := b.GetTile().SpawnUnit(b.Template.Unit)
	b.RegisterUnit(unit)
	_ = unit
}

func (b *Building) AddUnitSpawnProgress(value int) {
	b.Set(UnitSpawnProgress, b.Get(UnitSpawnProgress)+value)
}

func (b *Building) UnitSpawnIsComplete() bool {
	return b.Get(UnitSpawnProgress) >= b.Template.UnitTimer
}

func (b *Building) RegisterUnit(unit *Unit) {
	b.Units = append(b.Units, unit.GetPointer())
}

func (b *Building) GetName() string {
	return b.Name
}

func (b *Building) GetId() BuildingKey {
	return b.Id
}

func (b *Building) GetSector() *Sector {
	return b.GetTile().GetSector()
}

func (b *Building) GetFaction() *Faction {
	return b.GetSector().GetFaction()
}

func (b *Building) GetTile() *Tile {
	return theWorld.Grid.Tile(b.Location)
}

func (b *Building) GetDescription() string {
	stats := "Faction: " + b.GetFaction().GetName()
	if !b.ConstructionIsComplete() {
		stats += "\nConstruction Progress: " + strconv.Itoa(b.GetConstructionProgress())
		stats += " / " + strconv.Itoa(b.Template.ConstructionWork)
	}
	if b.Template.UnitLimit > 0 {
		currentUnits := len(b.Units)
		maxUnits := b.Template.UnitLimit
		housing := strconv.Itoa(currentUnits) + " / " + strconv.Itoa(maxUnits)
		stats += "\n\nUnit: " + b.Template.Unit
		stats += "\nCurrently housing " + housing
		if currentUnits < maxUnits {
			unitCost := b.Template.UnitTimer
			unitProgress := b.Get(UnitSpawnProgress)
			stats += "\nProgress to spawn: " + strconv.Itoa(unitProgress) + " / " + strconv.Itoa(unitCost)
		}
	}
	stats += "\n\n" + b.Template.Description
	return stats
}

func (b *Building) GetStats() string {
	stats := ""
	return stats
}

func (b *Building) GetPointer() BuildingPointer {
	return MakeBuildingPointer(b.GetId())
}

func (b *Building) Get(key int) int {
	return b.Attributes.Get(int(key))
}

func (b *Building) GetF(key int) float64 {
	return b.Attributes.GetF(key)
}

func (b *Building) Set(key, value int) {
	b.Attributes.Set(key, value)
}

func (b *Building) SetF(key int, value float64) {
	b.Attributes.SetF(key, value)
}
