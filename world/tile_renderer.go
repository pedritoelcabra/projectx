package world

import (
	"github.com/pedritoelcabra/projectx/gfx"
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
			RenderTile(x, y, RenderModeHeight, screen)
		}
	}
}

func RenderTile(x int, y int, mode TileRenderMode, screen *gfx.Screen) {
	return
}
