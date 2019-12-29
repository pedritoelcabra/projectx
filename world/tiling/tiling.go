package tiling

import (
	"math"
)

const (
	BaseTileSize = 36
	TileSize     = 36
	GfxTileSize  = 72
)

var TileWidth = 0.0
var TileHorizontalSeparation = 0.0
var TileWidthScale = 0.0
var TileHeight = 0.0
var TileHeightScale = 0.0
var Sqrt3 = 0.0
var One = 1

func InitTiling() {
	Sqrt3 = math.Sqrt(3.0)
	TileWidth = TileSize * 2
	TileWidthScale = TileWidth / GfxTileSize
	TileHeight = math.Floor(TileSize * Sqrt3)
	TileHeightScale = TileHeight / GfxTileSize
	TileHorizontalSeparation = math.Floor(TileWidth * 0.75)
}

func PixelFToTileI(x, y float64) (tx, ty int) {
	return PixelIToTileI(int(x), int(y))
}

func PixelIToTileI(x, y int) (tx, ty int) {
	coord := PixelFToTileC(float64(x), float64(y))
	return coord.X(), coord.Y()
}

func TileIToPixelF(tx, ty int) (x, y float64) {
	pixelCoord := TileCToPixelC(NewCoord(tx, ty))
	return float64(pixelCoord.X()), float64(pixelCoord.Y())
}

func PixelFToTileC(x, y float64) Coord {
	var q = (2.0 / 3.0 * x) / TileSize
	var r = (-1.0/3.0*x + Sqrt3/3.0*y) / TileSize
	return CubeToCoord(CubeRound(NewCube(q, -q-r, r)))
}

func TileCToPixelC(tileCoord Coord) Coord {
	var x = TileSize * (3 / 2) * float64(tileCoord.X())
	var y = TileSize * Sqrt3 * (float64(tileCoord.Y()) + 0.5*float64(tileCoord.X()&One))
	return NewCoordF(x, y)
}

func CubeRound(cube Cube) Cube {
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

	return NewCube(rx, ry, rz)
}

func CubeToCoord(cube Cube) Coord {
	var col = cube.X
	var row = cube.Z + (cube.X-float64(int(cube.X)&One))/2
	return NewCoord(int(col), int(row))
}
