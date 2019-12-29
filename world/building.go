package world

import (
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/world/grid"
)

type Building struct {
	Sprite    gfx.Sprite `json:"-"`
	SpriteKey gfx.SpriteKey
	Location  *grid.Tile
	X         float64
	Y         float64
}

func NewBuilding(spriteName string, location *grid.Tile) *Building {
	aBuilding := &Building{}
	aBuilding.SpriteKey = gfx.GetSpriteKey(spriteName)
	aBuilding.Sprite = gfx.NewHexSprite(aBuilding.SpriteKey)
	aBuilding.Location = location
	aBuilding.X = location.GetF(grid.RenderX)
	aBuilding.Y = location.GetF(grid.RenderY)
	return aBuilding
}

func (b *Building) DrawSprite(screen *gfx.Screen) {
	b.Sprite.DrawSprite(screen, b.X, b.Y)
}

func (b *Building) SetPosition(x, y float64) {

}

func (b *Building) Update(tick int, grid *grid.Grid) {

}
