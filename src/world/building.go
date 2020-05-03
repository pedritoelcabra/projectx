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
	Id                   BuildingKey
	Sprite               gfx.Sprite `json:"-"`
	SpriteKey            gfx.SpriteKey
	Name                 string
	Location             tiling.Coord
	X                    float64
	Y                    float64
	Template             *defs.BuildingDef
	ConstructionProgress int
}

func NewBuilding(name string, location *Tile) *Building {
	buildingDefs := defs.BuildingDefs()
	aBuilding := &Building{}
	aBuilding.Template = buildingDefs[name]
	aBuilding.ConstructionProgress = aBuilding.Template.ConstructionWork
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
	b.ConstructionProgress = value
}

func (b *Building) GetConstructionProgress() int {
	return b.ConstructionProgress
}

func (b *Building) AddConstructionProgress(value int) {
	b.SetConstructionProgress(value + b.GetConstructionProgress())
}

func (b *Building) ConstructionIsComplete() bool {
	return b.ConstructionProgress >= b.Template.ConstructionWork
}

func (b *Building) CompleteConstruction() {
	b.ConstructionProgress = b.Template.ConstructionWork
	b.SpriteKey = gfx.GetSpriteKey(b.Template.Graphic)
	b.Init()
}

func (b *Building) StartConstruction() {
	b.ConstructionProgress = 0
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

func (b *Building) Update(tick int, grid *Grid) {

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
	stats += "\n" + b.Template.Description
	return stats
}

func (b *Building) GetStats() string {
	stats := ""
	return stats
}

func (b *Building) GetPointer() BuildingPointer {
	return MakeBuildingPointer(b.GetId())
}
