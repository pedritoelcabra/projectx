package world

import (
	"github.com/pedritoelcabra/projectx/world/grid"
	"math"
)

const (
	BaseTileSize = 72
	TileSize     = 72
)

var TileHorizontalSeparation = 0
var TileScaleFactor = 0.0
var Sqrt3 = 0.0

func InitTiling() {
	TileHorizontalSeparation = int(TileSize * 0.75)
	TileScaleFactor = float64(TileSize) / float64(BaseTileSize)
	Sqrt3 = math.Sqrt(3.0)
}

func PosToTileC(coord grid.Coord) grid.Coord {
	return grid.NewCoord(PosToTile(coord.X(), coord.Y()))
}

func TileToPosC(tile grid.Coord) grid.Coord {
	return grid.NewCoord(PosToTile(tile.X(), tile.Y()))
}

func PosFloatToTile(x, y float64) (tx, ty int) {
	return PosToTile(int(x), int(y))
}

func PosToTile(x, y int) (tx, ty int) {
	coord := PixelToHex(float64(x), float64(y))
	return coord.X(), coord.Y()
}

func TileToPos(tx, ty int) (x, y int) {
	y = ty * TileSize
	x = tx * TileHorizontalSeparation
	if tx%2 != 0 {
		y += TileSize / 2
	}
	return
}

func TileToPosFloat(tx, ty int) (x, y float64) {
	ix, iy := TileToPos(tx, ty)
	return float64(ix), float64(iy)
}

func PixelToHex(x, y float64) grid.Coord {
	var q = (2.0 / 3.0 * x) / TileSize
	var r = (-1.0/3.0*x + Sqrt3/3.0*y) / TileSize
	return CubeToCoord(CubeRound(grid.NewCube(q, -q-r, r)))
}

func CubeRound(cube grid.Cube) grid.Cube {
	var rx = math.Round(cube.X)
	var ry = math.Round(cube.Y)
	var rz = math.Round(cube.Z)
	var x_diff = math.Abs(rx - cube.X)
	var y_diff = math.Abs(ry - cube.Y)
	var z_diff = math.Abs(rz - cube.Z)
	if x_diff > y_diff && x_diff > z_diff {
		rx = -ry - rz
	} else if y_diff > z_diff {
		ry = -rx - rz
	} else {
		rz = -rx - ry
	}

	return grid.NewCube(rx, ry, rz)
}

func CubeToCoord(cube grid.Cube) grid.Coord {
	var col = cube.X
	one := 1
	var row = cube.Z + (cube.X-float64(int(cube.X)&one))/2
	return grid.NewCoord(int(col), int(row))
}
