package world

import (
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/world/grid"
	"math"
)

type TileRenderMode int

const (
	RenderModeHeight TileRenderMode = iota
)

type TileRenderer struct {
	tileRenderSize int
}

func RenderTiles(screen *gfx.Screen, world *World) {
	tileSize := TileSize
	screenTileWidth := int(math.Ceil(float64(gfx.ScreenWidth / tileSize)))
	screenTileHeight := int(math.Ceil(float64(gfx.ScreenHeight / tileSize)))
	halfScreenTileWidth := int(math.Ceil(float64(screenTileWidth)))
	halfScreenTileHeight := int(math.Ceil(float64(screenTileHeight)))
	playerTileX, playerTileY := PosFloatToTile(world.PlayerUnit.GetPos())
	startX := playerTileX - halfScreenTileWidth
	endX := playerTileX + halfScreenTileWidth
	startY := playerTileY - halfScreenTileHeight
	endY := playerTileY + halfScreenTileHeight
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			RenderTile(world.Grid.Tile(grid.NewCoord(x, y)), RenderModeHeight, screen)
		}
	}
}

func RenderTile(tile *grid.tile, mode TileRenderMode, screen *gfx.Screen) {
	tx, ty := TileToPosFloat(x, y)
	gfx.DrawBasicTerrain(tx, ty, gfx.WaterFull, screen)
}
