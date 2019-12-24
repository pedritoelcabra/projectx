package grid

import "strconv"

// Coord is a simple structure to hold X and Y values
type Coord struct {
	x, y int
}

// Build a Coord from X and Y
func NewCoord(x, y int) Coord {
	aCoord := Coord{x, y}
	return aCoord
}

func (c Coord) X() int {
	return c.x
}

func (c Coord) Y() int {
	return c.y
}

func (c Coord) ToString() string {
	return strconv.Itoa(c.x) + "/" + strconv.Itoa(c.y)
}
