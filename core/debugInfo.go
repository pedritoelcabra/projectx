package core

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/world"
	"github.com/pedritoelcabra/projectx/src/world/tiling"
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
		//playerCoord := tiling.NewCoord(int(pX), int(pY))
		playerTileCoord := tiling.NewCoord(tiling.PixelIToTileI(int(pX), int(pY)))
		//aString += "\nPlayer Pos: " + playerCoord.ToString()
		//aString += "\nPlayer Tile: " + playerTileCoord.ToString()
		mx, my := ebiten.CursorPosition()
		cx, cy := g.Screen.GetCameraCoords()
		mx += int(cx)
		my += int(cy)
		mouseCoord := tiling.NewCoord(mx, my)
		aString += "\nMouse Pos: " + mouseCoord.ToString()
		mouseTileCoord := tiling.NewCoord(tiling.PixelIToTileI(mx, my))
		aString += "\nMouse Tile: " + mouseTileCoord.ToString()
		tile := g.World.Grid.Tile(mouseTileCoord)
		mHeight := tile.Get(world.Height)
		aString += "\nMouse Tile Height: " + strconv.Itoa(mHeight)
		aString += "\nMouse Tile Chunk: " + g.World.Grid.ChunkCoord(mouseTileCoord).ToString()

		playerMouseDist := tiling.HexDistance(playerTileCoord, mouseTileCoord)
		aString += "\nMouse distance to player: " + fmt.Sprintf("%f", playerMouseDist)

		aString += "\n-----"

		building := tile.GetBuilding()
		if building != nil {
			aString += "\nBuilding: " + building.GetName()
		}
		sector := g.World.GetSector(world.SectorKey(tile.Get(world.SectorId)))
		sectorName := "No mans land"
		if sector != nil {
			sectorName = sector.GetName()
		}
		aString += "\nSector: " + sectorName

	}
	if g.debugMessage != "" {
		aString += "\n" + g.debugMessage
	}

	return aString
}
