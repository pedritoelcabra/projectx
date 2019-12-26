package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/world/coord"
	"github.com/pedritoelcabra/projectx/world/grid"
	"github.com/pedritoelcabra/projectx/world/tiling"
	"image"
	"math"
)

type TileRenderMode int

const (
	RenderModeHeight TileRenderMode = iota
	RenderModeBasic
)

type TileRenderer struct {
	tileRenderSize int
}

func RenderTiles(screen *gfx.Screen, world *World, mode TileRenderMode) {
	tileSize := tiling.TileSize
	screenTileWidth := int(math.Ceil(float64(gfx.ScreenWidth / tileSize)))
	screenTileHeight := int(math.Ceil(float64(gfx.ScreenHeight / tileSize)))
	halfScreenTileWidth := int(math.Ceil(float64(screenTileWidth)))
	halfScreenTileHeight := int(math.Ceil(float64(screenTileHeight)))
	playerTileX, playerTileY := tiling.PixelFToTileI(world.PlayerUnit.GetPos())
	startX := playerTileX - halfScreenTileWidth
	endX := playerTileX + halfScreenTileWidth
	startY := playerTileY - halfScreenTileHeight
	endY := playerTileY + halfScreenTileHeight
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			RenderTile(world.Grid.Tile(coord.NewCoord(x, y)), mode, screen)
		}
	}
}

func RenderTile(tile *grid.Tile, mode TileRenderMode, screen *gfx.Screen) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(tiling.TileWidthScale, tiling.TileHeightScale)
	terrainBase := tile.Get(grid.TerrainBase)
	gfx.DrawHexTerrain(tile.GetF(grid.RenderX), tile.GetF(grid.RenderY), terrainBase, screen, op)
	//DrawDot(tile.GetF(grid.CenterX), tile.GetF(grid.CenterY), screen)
}

const (
	a0 = 0x40
	a1 = 0xc0
	a2 = 0xff
)

var pixels = []uint8{
	a0, a1, a1, a0,
	a1, a2, a2, a1,
	a1, a2, a2, a1,
	a0, a1, a1, a0,
}

var brushImage, _ = ebiten.NewImageFromImage(&image.Alpha{
	Pix:    pixels,
	Stride: 4,
	Rect:   image.Rect(0, 0, 4, 4),
}, ebiten.FilterDefault)

func DrawDot(x, y float64, screen *gfx.Screen) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x-2.0, y-2.0)
	screen.DrawImage(brushImage, op)
}
