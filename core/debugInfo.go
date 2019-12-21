package core

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pedritoelcabra/projectx/world"
	"github.com/pedritoelcabra/projectx/world/grid"
	"strconv"
)

func (g *game) DebugInfo() string {
	aString := "Tick: " + strconv.Itoa(g.tick)
	aString += "\nFrame: " + strconv.Itoa(g.framesDrawn)
	if g.World.PlayerUnit != nil {
		x, y := g.World.PlayerUnit.GetPos()
		aString += "\nPlayer Pos: " + strconv.Itoa(int(x)) + " / " + strconv.Itoa(int(y))
		tx, ty := world.PosToTile(int(x), int(y))
		aString += "\nPlayer Tile: " + strconv.Itoa(tx) + " / " + strconv.Itoa(ty)
		tile := g.World.Grid.Tile(grid.NewCoord(tx, ty))
		height := tile.Get(grid.Height)
		aString += "\nTile Height: " + strconv.Itoa(height)

		mx, my := ebiten.CursorPosition()
		cx, cy := g.Screen.GetCameraCoords()
		mx += int(cx)
		my += int(cy)
		aString += "\nMouse Pos: " + strconv.Itoa(int(mx)) + " / " + strconv.Itoa(int(my))
		mtx, mty := world.PosToTile(int(mx), int(my))
		aString += "\nMouse Tile: " + strconv.Itoa(mtx) + " / " + strconv.Itoa(mty)
		mTile := g.World.Grid.Tile(grid.NewCoord(mtx, mty))
		mHeight := mTile.Get(grid.Height)
		aString += "\nMouse Tile Height: " + strconv.Itoa(mHeight)
		aString += "\nMouse Tile Terrain: " + strconv.Itoa(mTile.Get(grid.TerrainBase))

	}
	if g.debugMessage != "" {
		aString += "\n" + g.debugMessage
	}

	return aString
}
