package grid

// Coord is a simple structure to hold X and Y values
type coord struct {
	x, y int
}

// Build a Coord from X and Y
func Coord(x, y int) coord {
	aCoord := coord{x, y}
	return aCoord
}

func (c coord) X() int {
	return c.x
}

func (c coord) Y() int {
	return c.y
}
