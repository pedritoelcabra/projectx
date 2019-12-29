package world

import (
	"github.com/pedritoelcabra/projectx/gfx"
)

type Building struct {
	ClassName string
	Sprite    gfx.Sprite `json:"-"`
	SpriteKey gfx.SpriteKey
	Location  *Tile
	X         float64
	Y         float64
}

func NewBuilding(spriteName string, location *Tile) *Building {
	aBuilding := &Building{}
	aBuilding.SpriteKey = gfx.GetSpriteKey(spriteName)
	aBuilding.Init()
	aBuilding.Location = location
	aBuilding.X = location.GetF(RenderX)
	aBuilding.Y = location.GetF(RenderY)
	return aBuilding
}

func (b *Building) Init() {
	b.ClassName = "Building"
	b.Sprite = gfx.NewHexSprite(b.SpriteKey)
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
