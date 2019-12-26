package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/world/coord"
	"github.com/pedritoelcabra/projectx/world/grid"
	"github.com/pedritoelcabra/projectx/world/tiling"
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
	op.GeoM.Scale(tiling.TileScaleFactor, tiling.TileScaleFactor)
	tx := tile.GetF(grid.RenderX)
	ty := tile.GetF(grid.RenderY)
	terrainBase := tile.Get(grid.TerrainBase)
	gfx.DrawHexTerrain(tx, ty, terrainBase, screen, op)
}
