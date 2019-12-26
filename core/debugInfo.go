package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/world"
	"github.com/pedritoelcabra/projectx/world/grid"
	"strconv"
)

func (g *game) DebugInfo() string {
	aString := ""
	aString += "\nFrame: " + strconv.Itoa(g.framesDrawn)
	if g.HasLoadedWorld() {
		aString += "\nTick: " + strconv.Itoa(g.World.GetTick())

		mx, my := ebiten.CursorPosition()
		cx, cy := g.Screen.GetCameraCoords()
		mx += int(cx)
		my += int(cy)
		aString += "\nMouse Pos: " + grid.NewCoord(mx, my).ToString()
		mouseTileCoord := grid.NewCoord(world.PixelToTile(mx, my))
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
