package world

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	tiling2 "github.com/pedritoelcabra/projectx/src/core/world/tiling"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/gui"
	"image"
	"strconv"
)

type BuildingKey int
type BuildingMap map[BuildingKey]*Building
type UnitList []UnitPointer

type Building struct {
	Id                     BuildingKey
	Sprite                 gfx.Sprite `json:"-"`
	ConstructionSprite     gfx.Sprite `json:"-"`
	SpriteKey              gfx.SpriteKey
	ConstructionSpriteKey  gfx.SpriteKey
	ConstructionPercentage float64
	Name                   string
	Location               tiling2.Coord
	X                      float64
	Y                      float64
	Template               *defs.BuildingDef
	Units                  UnitList
	Worker                 UnitPointer
	Attributes             *Attributes
}

func NewBuilding(name string, location *Tile) *Building {
	buildingDefs := defs.BuildingDefs()
	aBuilding := &Building{}
	aBuilding.Template = buildingDefs[name]
	aBuilding.Worker = MakeEmptyUnitPointer()
	aBuilding.Attributes = NewEmptyAttributes()
	aBuilding.Set(ConstructionProgress, aBuilding.Template.ConstructionWork)
	aBuilding.Set(UnitSpawnProgress, 0)
	aBuilding.SpriteKey = gfx.GetSpriteKey(aBuilding.Template.Graphic)
	aBuilding.ConstructionSpriteKey = gfx.GetSpriteKey(aBuilding.Template.ConstructionGraphic)
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
	b.ConstructionSprite = gfx.NewHexSprite(b.ConstructionSpriteKey)
	tile := theWorld.Grid.Tile(b.Location)
	if tile != nil {
		tile.SetBuilding(b)
	}
}

func (b *Building) SetConstructionProgress(value int) {
	if value < 0 {
		value = 0
	}
	b.ConstructionPercentage = (100.0 / float64(b.Template.ConstructionWork)) * float64(value)
	if value >= b.Template.ConstructionWork {
		value = b.Template.ConstructionWork
		b.CompleteConstruction()
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
	b.ConstructionPercentage = 100.0
	b.Init()
	b.FireWorker()
	sector := b.GetSector()
	if sector != nil {
		sector.GrowSectorToSize(b.Template.Influence, b.GetTile().GetCoord())
	}
}

func (b *Building) StartConstruction() {
	b.Set(ConstructionProgress, 0)
	b.ConstructionPercentage = 0
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
	if !b.ConstructionIsComplete() {
		b.ConstructionSprite.DrawSprite(screen, b.X, b.Y)
		b.Sprite.DrawSpriteSubImage(screen, b.X, b.Y, b.ConstructionPercentage)
		return
	}
	b.Sprite.DrawSprite(screen, b.X, b.Y)
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

func (b *Building) CheckForDeceasedUnits() {
	var aliveUnits = UnitList{}
	for _, unitPointer := range b.Units {
		unit := unitPointer.Get()
		if unit == nil {
			continue
		}
		if !unit.IsAlive() {
			continue
		}
		aliveUnits = append(aliveUnits, unitPointer)
	}
	b.Units = aliveUnits
}

func (b *Building) SpawnUnit() {
	b.Set(UnitSpawnProgress, 0)
	unit := b.GetTile().SpawnUnit(b.Template.Unit)
	unit.SetFaction(b.GetFaction())
	b.RegisterUnit(unit)
	unit.SetHome(b)
	_ = unit
}

func (b *Building) SpawnAllUnits() {
	limit := b.Template.UnitLimit
	for i := len(b.Units); i < limit; i++ {
		b.SpawnUnit()
	}
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

func (b *Building) GetStats() string {
	stats := "Faction: " + b.GetFaction().GetName()

	if !b.ConstructionIsComplete() {
		stats += "\n\nConstruction Progress: " + strconv.Itoa(b.GetConstructionProgress())
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

func (b *Building) HasWorkSlot() bool {
	currentWorker := b.Worker.Get()
	if currentWorker != nil {
		return false
	}
	if !b.ConstructionIsComplete() {
		return true
	}
	return false
}

func (b *Building) GetWorker() *Unit {
	return b.Worker.Get()
}

func (b *Building) SetWorker(unit *Unit) {
	b.Worker = unit.GetPointer()
}

func (b *Building) FireWorker() {
	worker := b.Worker.Get()
	if worker == nil {
		return
	}
	worker.WorkPlace = MakeEmptyBuildingPointer()
	b.Worker = MakeEmptyUnitPointer()
}

func (b *Building) AddWork() {
	if !b.ConstructionIsComplete() {
		b.AddConstructionProgress(1)
	}
}

func (b *Building) AddButtonsToEntityMenu(menu *gui.Menu, size image.Rectangle) {
	worker := b.Worker.Get()
	if worker != nil {
		workerButton := gui.NewButton(size, "Worker: "+worker.GetName())
		workerButton.OnPressed = func(b *gui.Button) {
			theWorld.SetDisplayEntity(worker)
		}
		menu.AddButton(workerButton)
	}
	for _, unitPointer := range b.Units {
		unit := unitPointer.Get()
		if unit == nil {
			continue
		}
		unitButton := gui.NewButton(size, "Inhabitant: "+unit.GetName())
		unitButton.OnPressed = func(b *gui.Button) {
			theWorld.SetDisplayEntity(unit)
		}
		menu.AddButton(unitButton)
	}
	sector := b.GetSector()
	if sector != nil {
		sectorButton := gui.NewButton(size, "Sector: "+sector.GetName())
		sectorButton.OnPressed = func(b *gui.Button) {
			theWorld.SetDisplayEntity(sector)
		}
		menu.AddButton(sectorButton)
	}
}
