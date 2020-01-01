package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/gfx"
	"github.com/pedritoelcabra/projectx/world/tiling"
	"math"
)

type TileRenderMode int

const (
	RenderModeHeight TileRenderMode = iota
	RenderModeBasic
)

var lastRenderedChunkCoord = tiling.NewCoord(-999, 999)

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
			RenderChunk(x, y, screen, world)
		}
	}
}

func RenderChunk(x, y int, screen *gfx.Screen, world *World) {
	chunkCoord := world.Grid.ChunkCoord(tiling.NewCoord(x, y))
	if lastRenderedChunkCoord != chunkCoord {
		lastRenderedChunkCoord = chunkCoord
		aChunk := world.Grid.Chunk(chunkCoord)
		aChunk.GenerateImage()
		chunkImage := aChunk.GetImage()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(aChunk.FirstTile().GetF(RenderX), aChunk.FirstTile().GetF(RenderY))
		screen.DrawImage(chunkImage, op)
		aChunk.RunOnAllTiles(DrawSectorBorders)
	}
}
