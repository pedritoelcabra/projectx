package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/world/grid"
	"log"
	"math"
	"strconv"
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

func RenderTile(tile *grid.Tile, mode TileRenderMode, screen *gfx.Screen) {
	op := &ebiten.DrawImageOptions{}
	tx, ty := TileToPosFloat(tile.X(), tile.Y())
	if mode == RenderModeHeight {
		RenderHeightMapTile(tile, screen, op)
		return
	}
	terrainBase := tile.Get(grid.TerrainBase)
	if terrainBase >= 0 {
		gfx.DrawBasicTerrain(tx, ty, gfx.BasicTerrainTypes(terrainBase), screen, op)
	}
}

func RenderHeightMapTile(tile *grid.Tile, screen *gfx.Screen, op *ebiten.DrawImageOptions) {
	tx, ty := TileToPosFloat(tile.X(), tile.Y())
	height := tile.Get(grid.Height)
	if height < -1000 || height > 1000 {
		log.Fatal("Unsupported height " + strconv.Itoa(height))
	}
	alpha := (float64(height) + 1000.0) / 2000.0
	op.ColorM.Scale(1.0, 1.0, 1.0, alpha)
	gfx.DrawBasicTerrain(tx, ty, gfx.BasicTerrainTypes(gfx.StoneBlockFull), screen, op)
}
