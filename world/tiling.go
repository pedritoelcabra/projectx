package world

import (
	"github.com/pedritoelcabra/projectx/world/grid"
)

const (
	TileSize = 32
)

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
	tx = int(x / TileSize)
	if x < 0 {
		tx--
	}
	ty = int(y / TileSize)
	if y < 0 {
		ty--
	}
	return
}

func TileToPos(tx, ty int) (x, y int) {
	x = tx * TileSize
	y = ty * TileSize
	return
}

func TileToPosFloat(tx, ty int) (x, y float64) {
	ix, iy := TileToPos(tx, ty)
	return float64(ix), float64(iy)
}
