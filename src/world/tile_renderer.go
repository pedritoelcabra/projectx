package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/gfx"
	tiling2 "github.com/pedritoelcabra/projectx/src/world/tiling"
	"math"
)

type TileRenderMode int

const (
	RenderModeHeight TileRenderMode = iota
	RenderModeBasic
)

var lastRenderedChunkCoord = tiling2.NewCoord(-999, 999)

var tilesToRender []*Tile
var playerLastCoord tiling2.Coord

func InitTileRenderer() {
	playerLastCoord = tiling2.NewCoord(-99999, -99999)
	tilesToRender = []*Tile{}
}

func RenderTiles(screen *gfx.Screen, world *World, mode TileRenderMode) {
	LoadTilesToRender(world)
	CallTileRenderFunction(DrawTerrain)
	CallTileRenderFunction(DrawSectorBorders)
}

func CallTileRenderFunction(drawFun func(*Tile)) {
	for _, tile := range tilesToRender {
		drawFun(tile)
	}
}

func RenderChunk(x, y int, screen *gfx.Screen, world *World) {
	chunkCoord := world.Grid.ChunkCoord(tiling2.NewCoord(x, y))
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

func LoadTilesToRender(world *World) {
	playerCoord := tiling2.NewCoord(tiling2.PixelFToTileI(world.PlayerUnit.GetPos()))
	if playerCoord.Equals(playerLastCoord) {
		return
	}
	playerLastCoord = playerCoord
	tilesToRender = []*Tile{}

	screenTileWidth := int(math.Ceil(float64(gfx.ScreenWidth / tiling2.TileHorizontalSeparation)))
	screenTileHeight := int(math.Ceil(float64(gfx.ScreenHeight / tiling2.TileHeight)))
	halfScreenTileWidth := int(math.Ceil(float64(screenTileWidth / 2)))
	halfScreenTileHeight := int(math.Ceil(float64(screenTileHeight / 2)))
	playerTileX, playerTileY := tiling2.PixelFToTileI(world.PlayerUnit.GetPos())
	startX := playerTileX - halfScreenTileWidth - 1
	endX := playerTileX + halfScreenTileWidth + 1
	startY := playerTileY - halfScreenTileHeight - 1
	endY := playerTileY + halfScreenTileHeight + 1
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			tilesToRender = append(tilesToRender, world.Grid.Tile(tiling2.NewCoord(x, y)))
		}
	}
}
