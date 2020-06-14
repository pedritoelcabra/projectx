package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/src/core/defs"
	"github.com/pedritoelcabra/projectx/src/core/world"
	"github.com/pedritoelcabra/projectx/src/core/world/tiling"
	"strconv"
)

func (g *game) DebugInfo() string {

	displayResourceTotals := true

	aString := ""
	aString += "\nFPS: " + strconv.Itoa(int(ebiten.CurrentFPS()))
	aString += "\nTPS: " + strconv.Itoa(int(ebiten.CurrentTPS()))
	//aString += "\nFrame: " + strconv.Itoa(g.framesDrawn)
	if g.HasLoadedWorld() {
		//aString += "\nTick: " + strconv.Itoa(g.World.GetTick())

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
		mouseTileCoord := g.MouseTileCoord()
		aString += "\nMouse Tile: " + mouseTileCoord.ToString()
		tile := g.World.Grid.Tile(mouseTileCoord)
		mHeight := tile.Get(world.Height)
		aString += "\nMouse Tile Height: " + strconv.Itoa(mHeight)
		aString += "\nMouse Tile Chunk: " + g.World.Grid.ChunkCoord(mouseTileCoord).ToString()

		//playerMouseDist := tiling.HexDistance(playerTileCoord, mouseTileCoord)
		//aString += "\nMouse distance to player: " + fmt.Sprintf("%f", playerMouseDist)

		//drawUnitCount := g.World.Data.Get(world.CurrentDrawnUnits)
		//drawBuildingCount := g.World.Data.Get(world.CurrentDrawnBuildings)
		//aString += "\nDrawing " + strconv.Itoa(drawUnitCount) + " units, " + strconv.Itoa(drawBuildingCount) + " buildings"
		//aString += "\n-----"

		//building := tile.GetBuilding()
		//if building != nil {
		//	aString += "\nBuilding: " + building.GetName()
		//}
		//veg := tile.Get(world.Resource)
		//if veg != 0 {
		//	aString += "\n" + defs.ResourceById(veg).Name
		//}
		//sector := g.World.GetSector(world.SectorKey(tile.Get(world.SectorId)))
		//sectorName := "No mans land"
		//if sector != nil {
		//	factionName := g.World.GetFaction(world.FactionKey(sector.Get(world.FactionId))).GetName()
		//	sectorName = sector.GetName() + " (" + factionName + ")"
		//}
		//aString += "\nSector: " + sectorName

		if displayResourceTotals {
			aString += "\nTotal spawned resource tiles: "
			matDefs := defs.GetMaterialDefs()
			for i := 1; i <= len(matDefs); i++ {
				matDef := defs.GetMaterialDefByKey(i)
				aString += "\n" + matDef.Name + ": " + strconv.Itoa(defs.GetResourceLocationTotals(matDef.Name))
			}
		}

		nearestSectorKey := world.SectorKey(-1)
		nearestSectorDist := 999999
		for _, nearSector := range g.World.Entities.Sectors {
			dist := nearSector.GetCenter().ChebyshevDist(playerTileCoord)
			if dist < nearestSectorDist {
				nearestSectorDist = dist
				nearestSectorKey = nearSector.GetId()
			}
		}
		if nearestSectorKey != -1 {
			nearestSector := g.World.GetSector(nearestSectorKey)
			xDist := nearestSector.GetCenter().X() - playerTileCoord.X()
			yDist := nearestSector.GetCenter().Y() - playerTileCoord.Y()
			aString += "\n\nNearest sector to Player: " + nearestSector.GetName()
			aString += "\n" + strconv.Itoa(xDist) + " / " + strconv.Itoa(yDist)
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

func (g *game) CurrentMouseTile() *world.Tile {
	pos := g.MouseTileCoord()
	return g.World.Grid.Tile(pos)
}
