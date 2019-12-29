package tiling

import (
	"github.com/pedritoelcabra/projectx/world/utils"
	"strconv"
)

type Coord struct {
	CX int `json:"X"`
	CY int `json:"Y"`
}

type Cube struct {
	X float64
	Y float64
	Z float64
}

func NewCoord(x, y int) Coord {
	aCoord := Coord{x, y}
	return aCoord
}

func NewCube(x, y, z float64) Cube {
	aCube := Cube{x, y, z}
	return aCube
}

func NewCoordF(x, y float64) Coord {
	return NewCoord(int(x), int(y))
}

func (c Coord) X() int {
	return c.CX
}

func (c Coord) Y() int {
	return c.CY
}

func (c Coord) ToString() string {
	return strconv.Itoa(c.CX) + "/" + strconv.Itoa(c.CY)
}

func (c1 Coord) ManhattanDist(c2 Coord) int {
	yDist := utils.AbsInt(c1.CY - c2.CY)
	xDist := utils.AbsInt(c1.CX - c2.CX)
	return yDist + xDist
}

func (c1 Coord) ChebyshevDist(c2 Coord) int {
	yDist := utils.AbsInt(c1.CY - c2.CY)
	xDist := utils.AbsInt(c1.CX - c2.CX)
	return utils.MaxInt(yDist, xDist)
}
