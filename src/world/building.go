package world

import (
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/gfx"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
)

type BuildingKey int
type BuildingMap map[BuildingKey]*Building

type Building struct {
	Id        EntityKey
	ClassName string
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
	aBuilding.Id = theWorld.AddEntity(aBuilding)
	location.SetBuilding(aBuilding)
	return aBuilding
}

func (b *Building) Init() {
	b.ClassName = "Building"
	b.Sprite = gfx.NewHexSprite(b.SpriteKey)
	tile := theWorld.Grid.Tile(b.Location)
	if tile != nil {
		tile.SetBuilding(b)
		b.X = tile.GetF(RenderX)
		b.Y = tile.GetF(RenderY)
	}
}

func (b *Building) DrawSprite(screen *gfx.Screen) {
	b.Sprite.DrawSprite(screen, b.X, b.Y)
}

func (b *Building) SetPosition(x, y float64) {

}

func (b *Building) Update(tick int, grid *Grid) {

}

func (b *Building) GetClassName() string {
	return b.ClassName
}

func (b *Building) GetName() string {
	return b.Name
}

func (b *Building) GetId() EntityKey {
	return b.Id
}
