// A package for managing a Grid based game world
package grid

import "log"

type Grid struct {
	size   int
	radius int
	tiles  []tile
}

func New(size int) Grid {
	if size < 0 {
		size = -size
	}
	if size%2 == 1 {
		size++
	}
	arraySize := size * size
	arrayTiles := make([]tile, arraySize)
	return Grid{size, size / 2, arrayTiles}
}

func (g Grid) Size() int {
	return g.size
}

func (g Grid) Radius() int {
	return g.radius
}

func (g Grid) Tile(c coord) tile {
	x, y := c.X(), c.Y()
	if x < 1 || x > g.Size() || y < 1 || y > g.Size() {
		log.Fatalf("Grid.Tile() requested invalid coordinates %d / %d", x, y)
	}
	aTile := g.tiles[g.gridIndex(x, y)]
	if aTile.X() == 0 {
		g.tiles[g.gridIndex(x, y)] = g.initTile(c)
		aTile = g.tiles[g.gridIndex(x, y)]
	}
	return aTile
}

func (g Grid) initTile(c coord) tile {
	return tile{c, make(map[int]int)}
}

func (g Grid) gridIndex(x, y int) int {
	return (x * g.size) + y
}

func (g Grid) gridCoordinates(index int) (x, y int) {
	y = index % g.size
	x = (index - y) / g.size
	return
}
