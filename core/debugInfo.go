package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/world"
	"github.com/pedritoelcabra/projectx/world/tiling"
	"strconv"
	"time"
)

var lastSecond = 0
var lastFramesDrawn = 0
var lastFPS = 0

func (g *game) DebugInfo() string {
	second := int(time.Now().Unix())
	if lastSecond != second {
		lastSecond = second
		lastFPS = g.framesDrawn - lastFramesDrawn
		lastFramesDrawn = g.framesDrawn
	}
	aString := ""
	aString += "\nFPS: " + strconv.Itoa(lastFPS)
	aString += "\nFrame: " + strconv.Itoa(g.framesDrawn)
	if g.HasLoadedWorld() {
		aString += "\nTick: " + strconv.Itoa(g.World.GetTick())

		pX, pY := g.World.PlayerUnit.GetPos()
		aString += "\nPlayer Pos: " + tiling.NewCoord(int(pX), int(pY)).ToString()
		mx, my := ebiten.CursorPosition()
		cx, cy := g.Screen.GetCameraCoords()
		mx += int(cx)
		my += int(cy)
		aString += "\nMouse Pos: " + tiling.NewCoord(mx, my).ToString()
		mouseTileCoord := tiling.NewCoord(tiling.PixelIToTileI(mx, my))
		aString += "\nMouse Tile: " + mouseTileCoord.ToString()
		mHeight := g.World.Grid.Tile(mouseTileCoord).Get(world.Height)
		aString += "\nMouse Tile Height: " + strconv.Itoa(mHeight)
		aString += "\nMouse Tile Chunk: " + g.World.Grid.ChunkCoord(mouseTileCoord).ToString()

	}
	if g.debugMessage != "" {
		aString += "\n" + g.debugMessage
	}

	return aString
}
