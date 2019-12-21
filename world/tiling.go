package world

import (
	"github.com/pedritoelcabra/projectx/world/grid"
)

const (
	BaseTileSize = 72
	TileSize     = 72
)

var TileHorizontalSeparation = 0
var TileScaleFactor = 0.0

func InitTiling() {
	TileHorizontalSeparation = int(TileSize * 0.75)
	TileScaleFactor = float64(TileSize) / float64(BaseTileSize)
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
	tx = x / TileHorizontalSeparation
	if x < 0 {
		tx--
	}
	if tx%2 != 0 {
		y += TileSize / 2
	}
	ty = y / TileSize
	if y < 0 {
		ty--
	}
	return
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
