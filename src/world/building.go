package world

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
)

type BuildingKey int
type BuildingMap map[BuildingKey]*Building

type Building struct {
	Id        BuildingKey
	Sprite    gfx.Sprite `json:"-"`
	SpriteKey gfx.SpriteKey
	Name      string
	Location  tiling.Coord
	X         float64
	Y         float64
	Template  *defs.BuildingDef
}

func NewBuilding(name string, location *Tile) *Building {
	buildingDefs := defs.BuildingDefs()
	aBuilding := &Building{}
	aBuilding.Template = buildingDefs[name]
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
	return b.Template.Description
}
