package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/core/defs"
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

		//pX, pY := g.World.PlayerUnit.GetPos()
		//playerCoord := tiling.NewCoord(int(pX), int(pY))
		//playerTileCoord := tiling.NewCoord(tiling.PixelIToTileI(int(pX), int(pY)))
		//aString += "\nPlayer Pos: " + playerCoord.ToString()
		//aString += "\nPlayer Tile: " + playerTileCoord.ToString()
		mx, my := ebiten.CursorPosition()
		cx, cy := g.Screen.GetCameraCoords()
		mx += int(cx)
		my += int(cy)
		mouseCoord := tiling.NewCoord(mx, my)
		aString += "\nMouse Pos: " + mouseCoord.ToString()
		mouseTileCoord := g.MouseTileCoord()
		aString += "\nMouse Tile: " + mouseTileCoord.ToString()
		tile := g.World.Grid.Tile(mouseTileCoord)
		mHeight := tile.Get(world.Height)
		aString += "\nMouse Tile Height: " + strconv.Itoa(mHeight)
		aString += "\nMouse Tile Chunk: " + g.World.Grid.ChunkCoord(mouseTileCoord).ToString()

		//playerMouseDist := tiling.HexDistance(playerTileCoord, mouseTileCoord)
		//aString += "\nMouse distance to player: " + fmt.Sprintf("%f", playerMouseDist)

		drawEntityCount := g.World.Data.Get(world.CurrentDrawnEntities)
		aString += "\nCurrently drawing " + strconv.Itoa(drawEntityCount) + " units"
		aString += "\n-----"

		building := tile.GetBuilding()
		if building != nil {
			aString += "\nBuilding: " + building.GetName()
		}
		veg := tile.Get(world.Flora)
		if veg != 0 {
			aString += "\n" + defs.VegetationById(veg).Name
		}
		sector := g.World.GetSector(world.SectorKey(tile.Get(world.SectorId)))
		sectorName := "No mans land"
		if sector != nil {
			sectorName = sector.GetName() + " owned by " + g.World.GetFaction(world.FactionKey(sector.Get(world.FactionId))).GetName()
		}
		aString += "\nSector: " + sectorName

		unitsAtLocation := g.World.UnitsCollidingWith(float64(mouseCoord.X()), float64(mouseCoord.Y()))
		for _, unit := range unitsAtLocation {
			aString += "\n" + unit.GetName()
			playerFaction := unit.GetFaction()
			if playerFaction != nil {
				aString += " (" + playerFaction.GetName() + ")"
			}
		}
	}
	if g.debugMessage != "" {
		aString += "\n" + g.debugMessage
	}

	return aString
}

func (g *game) MouseTileCoord() tiling.Coord {
	mx, my := ebiten.CursorPosition()
	cx, cy := g.Screen.GetCameraCoords()
	mx += int(cx)
	my += int(cy)
	return tiling.NewCoord(tiling.PixelIToTileI(mx, my))
}

func (g *game) MousePosCoord() tiling.Coord {
	mx, my := ebiten.CursorPosition()
	cx, cy := g.Screen.GetCameraCoords()
	mx += int(cx)
	my += int(cy)
	return tiling.NewCoord(mx, my)
}
