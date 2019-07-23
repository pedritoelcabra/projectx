package grid

import (
	"testing"
)

func TestSize(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{12, 12},
		{123, 124},
		{-12, 12},
		{-11, 12},
		{0, 0},
	}
	for _, c := range cases {
		got := New(c.in)
		radius := got.Size()
		if radius != c.want {
			t.Errorf("%d == %d, want %d", c.in, radius, c.want)
		}
	}
}

func TestRadius(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{12, 6},
		{123, 62},
		{-12, 6},
		{-11, 6},
		{0, 0},
	}
	for _, c := range cases {
		got := New(c.in)
		radius := got.Radius()
		if radius != c.want {
			t.Errorf("%d == %d, want %d", c.in, radius, c.want)
		}
	}
}

func TestGrid_Tile(t *testing.T) {
	grid1 := New(100)
	tile1 := grid1.Tile(Coord(1, 1))
	if tile1.X() != 1 || tile1.Y() != 1 {
		t.Errorf("Tile at %d / %d, returned %d / %d", 1, 1, tile1.X(), tile1.Y())
	}
	tile2 := grid1.Tile(Coord(23, 47))
	if tile2.X() != 23 || tile2.Y() != 47 {
		t.Errorf("Tile at %d / %d, returned %d / %d", 1, 1, tile2.X(), tile2.Y())
	}
}

func TestGrid_Conversions(t *testing.T) {
	cases := []struct {
		x, y int
	}{
		{1, 1},
		{12, 99},
		{43, 60},
		{22, 1},
	}
	aGrid := New(100)
	for _, c := range cases {
		index := aGrid.gridIndex(c.x, c.y)
		x, y := aGrid.gridCoordinates(index)
		if x != c.x || y != c.y {
			t.Errorf("%d / %d converted to %d / %d", c.x, c.y, x, y)
		}
	}
}

func TestGrid_SetGet(t *testing.T) {
	const Height = 1
	aGrid := New(100)
	gotHeight := aGrid.Tile(Coord(1, 1)).Get(Height)
	if gotHeight != 0 {
		t.Errorf("Expected uninitialized value to be %d, got %d", 0, gotHeight)
	}
	newHeight := 55
	aGrid.Tile(Coord(1, 1)).Set(Height, newHeight)
	gotHeight = aGrid.Tile(Coord(1, 1)).Get(Height)
	if gotHeight != newHeight {
		t.Errorf("Expected value to be %d, got %d", newHeight, gotHeight)
	}
	gotHeight = aGrid.Tile(Coord(2, 1)).Get(Height)
	if gotHeight != 0 {
		t.Errorf("Expected uninitialized value to be %d, got %d", 0, gotHeight)
	}
}
