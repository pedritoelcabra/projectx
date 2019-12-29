package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/world/coord"
	"github.com/pedritoelcabra/projectx/world/grid"
	"github.com/pedritoelcabra/projectx/world/tiling"
	"strconv"
)

func (g *game) DebugInfo() string {
	aString := ""
	aString += "\nFrame: " + strconv.Itoa(g.framesDrawn)
	if g.HasLoadedWorld() {
		aString += "\nTick: " + strconv.Itoa(g.World.GetTick())

		pX, pY := g.World.PlayerUnit.GetPos()
		aString += "\nPlayer Pos: " + coord.NewCoord(int(pX), int(pY)).ToString()
		mx, my := ebiten.CursorPosition()
		cx, cy := g.Screen.GetCameraCoords()
		mx += int(cx)
		my += int(cy)
		aString += "\nMouse Pos: " + coord.NewCoord(mx, my).ToString()
		mouseTileCoord := coord.NewCoord(tiling.PixelIToTileI(mx, my))
		aString += "\nMouse Tile: " + mouseTileCoord.ToString()
		mHeight := g.World.Grid.Tile(mouseTileCoord).Get(grid.Height)
		aString += "\nMouse Tile Height: " + strconv.Itoa(mHeight)
		aString += "\nMouse Tile Chunk: " + g.World.Grid.ChunkCoord(mouseTileCoord).ToString()

	}
	if g.debugMessage != "" {
		aString += "\n" + g.debugMessage
	}

	return aString
}
