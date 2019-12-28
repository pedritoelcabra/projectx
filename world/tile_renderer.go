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

var lastRenderedChunkCoord = coord.NewCoord(-999, 999)

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
	chunkCoord := world.Grid.ChunkCoord(coord.NewCoord(x, y))
	if lastRenderedChunkCoord == chunkCoord {
		return
	}
	lastRenderedChunkCoord = chunkCoord
	aChunk := world.Grid.Chunk(chunkCoord)
	aChunk.GenerateImage()
	chunkImage := aChunk.GetImage()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(aChunk.FirstTile().GetF(grid.RenderX), aChunk.FirstTile().GetF(grid.RenderY))
	screen.DrawImage(chunkImage, op)
	if aChunk.Sector != nil {
		sectorCoord := aChunk.Sector.Center
		sectorTile := world.Grid.Tile(sectorCoord)
		DrawDot(sectorTile.GetF(grid.CenterX), sectorTile.GetF(grid.CenterY), screen, 4.0)
	}
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

func DrawDot(x, y float64, screen *gfx.Screen, scale float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x-8.0, y-8.0)
	screen.DrawImage(brushImage, op)
}
