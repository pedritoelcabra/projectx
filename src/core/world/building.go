package world

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/inventory"
	"github.com/pedritoelcabra/projectx/src/core/world/tiling"
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
	Location               tiling.Coord
	X                      float64
	Y                      float64
	Template               *defs.BuildingDef
	Units                  UnitList
	Attributes             *Attributes
	Job                    *Job
	ResourceRequests       inventory.RequestList
}

func NewBuilding(name string, location *Tile) *Building {
	buildingDefs := defs.BuildingDefs()
	aBuilding := &Building{}
	aBuilding.Template = buildingDefs[name]
	aBuilding.ResourceRequests = make([]*inventory.ResourceRequest, 0)
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
	b.DestroyJob()
	sector := b.GetSector()
	if sector != nil {
		sector.GrowSectorToSize(b.Template.Influence, b.GetTile().GetCoord())
	}
}

func (b *Building) StartConstruction() {
	b.Set(ConstructionProgress, 0)
	b.ConstructionPercentage = 0
	b.LoadResourceRequests(b.Template.Resources)
	b.FulfillResourceRequests()
	b.Init()
}

func (b *Building) LoadResourceRequests(source []*defs.ResourceRequirement) {
	b.ResourceRequests = inventory.CopyResourceRequirements(source, b.ResourceRequests)
}

func (b *Building) FulfillResourceRequests() {
	b.ResourceRequests = inventory.FulfillResourceRequests(b.GetSector().Inventory, b.ResourceRequests)
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
	b.UpdateJobSpawn()
}

func (b *Building) UpdateJobSpawn() {
	if b.Job != nil {
		return
	}
	if len(b.ResourceRequests) != 0 {
		return
	}
	if !b.ConstructionIsComplete() {
		b.Job = NewBuildJob(b)
	}
	if b.Job != nil {
		b.GetSector().PublishJob(b.Job)
		return
	}
	if b.Template.Gathers != "" {
		b.Job = b.SearchGatherJob()
	}
	if b.Job != nil {
		b.GetSector().PublishJob(b.Job)
		return
	}
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

	if len(b.ResourceRequests) > 0 {
		stats += "\n\nNeeded resources:"
		for _, request := range b.ResourceRequests {
			def := defs.GetMaterialDefByKey(request.Type)
			stats += "\n" + def.Name + ": " + strconv.Itoa(request.Amount)
		}
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

func (b *Building) SearchGatherJob() *Job {
	if b.Template.Gathers == "" {
		return nil
	}
	status := b.Get(GatherStatus)
	if status == GatherStatusExhausted {
		return nil
	}
	b.Set(GatherStatus, GatherStatusExhausted)
	for _, t := range b.GetTile().TilesInRadius(b.Template.GatherRadius) {
		if t.GetSector() == nil || b.GetSector() == nil || t.GetSector().GetId() != b.GetSector().GetId() {
			continue
		}
		if !t.HasResource(b.Template.Gathers) {
			continue
		}
		b.Set(GatherStatus, GatherStatusAvailable)
		return NewGatheringJob(b, t.GetCoord())
	}
	return nil
}

func (b *Building) GetWorker() *Unit {
	if b.Job != nil {
		return b.Job.GetWorker()
	}
	return nil
}

func (b *Building) ClearJob() {
	b.Job = nil
}

func (b *Building) DestroyJob() {
	if b.Job == nil {
		return
	}
	b.Job.Destroy()
}

func (b *Building) GetWorkLocation() tiling.Coord {
	if !b.ConstructionIsComplete() {
		return b.GetTile().GetCoord()
	}
	if b.Get(GatherStatus) == GatherStatusAvailable {
		return tiling.NewCoord(b.Get(GatherTargetX), b.Get(GatherTargetY))
	}
	return b.GetTile().GetCoord()
}

func (b *Building) AddButtonsToEntityMenu(menu *gui.Menu, size image.Rectangle) {
	if b.Job != nil {
		worker := b.Job.GetWorker()
		if worker != nil {
			workerButton := gui.NewButton(size, "Worker: "+worker.GetName())
			workerButton.OnPressed = func(b *gui.Button) {
				theWorld.SetDisplayEntity(worker)
			}
			menu.AddButton(workerButton)
		} else {
			workerButton := gui.NewButton(size, b.Job.GetName())
			menu.AddButton(workerButton)
		}
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
