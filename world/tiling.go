package world

import (
	"github.com/pedritoelcabra/projectx/world/grid"
)

const (
	BaseTileSize = 32
	TileSize     = 32
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
	ty = y / TileSize
	if y < 0 {
		ty--
	}
	if ty%2 > 0 {
		x += TileSize / 2
	}
	tx = x / TileSize
	if x < 0 {
		tx--
	}
	return
}

func TileToPos(tx, ty int) (x, y int) {
	y = ty * TileSize
	x = tx * TileSize
	if ty%2 > 0 {
		x += TileSize / 2
	}
	return
}

func TileToPosFloat(tx, ty int) (x, y float64) {
	ix, iy := TileToPos(tx, ty)
	return float64(ix), float64(iy)
}
