package grid

import (
	"testing"
)

func TestChunkCoord(t *testing.T) {
	cases := []struct {
		tx, ty, cx, cy int
	}{
		{0, 0, 0, 0},
		{31, 31, 0, 0},
		{32, 32, 1, 1},
		{32, -1, 1, -1},
		{-1, -1, -1, -1},
		{-32, -32, -1, -1},
		{-33, -32, -2, -1},
		{TileOffset - 32, TileOffset - 32, GridOffset - 1, GridOffset - 1},
	}
	for _, c := range cases {
		grid := New()
		chunkCoord := grid.chunkCoord(Coord(c.tx, c.ty))
		if chunkCoord.X() != c.cx || chunkCoord.Y() != c.cy {
			t.Errorf("%d / %d, want %d / %d", chunkCoord.X(), chunkCoord.Y(), c.cx, c.cy)
		}
	}
}

func TestChunkIndex(t *testing.T) {
	cases := []struct {
		cx, cy, ci int
	}{
		{0, 0, 500500},
		{31, 31, 531531},
		{32, 32, 532532},
		{32, -1, 532499},
		{-1, -1, 499499},
		{-32, -32, 468468},
		{-33, -32, 467468},
	}
	for _, c := range cases {
		grid := New()
		chunkIndex := grid.chunkIndex(c.cx, c.cy)
		if chunkIndex != c.ci {
			t.Errorf("%d, want %d", chunkIndex, c.ci)
		}
	}
}

func TestTileCoord(t *testing.T) {
	cases := []struct {
		x, y int
	}{
		{0, 0},
		{100, -100},
		{-1500, 1200},
		{256, 256},
	}
	for _, c := range cases {
		grid := New()
		aTile := grid.Tile(Coord(c.x, c.y))
		if aTile.X() != c.x || aTile.Y() != c.y {
			t.Errorf("%d / %d, want %d / %d", aTile.X(), aTile.Y(), c.x, c.y)
		}
	}
}

func TestTileVariables(t *testing.T) {

	aGrid := New()
	const Population = 1
	const Height = 2

	aCoord := Coord(100, 200)
	bCoord := Coord(0, 0)
	aTile := aGrid.Tile(aCoord)
	bTile := aGrid.Tile(bCoord)

	setValue := 10000
	aTile.Set(Population, setValue)
	gotValue := aTile.Get(Population)
	if setValue != gotValue {
		t.Errorf("Set %d - Got %d", setValue, gotValue)
	}

	setValue = 500
	aTile.Set(Height, setValue)
	gotValue = aTile.Get(Height)
	if setValue != gotValue {
		t.Errorf("Set %d - Got %d", setValue, gotValue)
	}

	setValue = 2222
	aTile.Set(Population, setValue)
	gotValue = aTile.Get(Population)
	if setValue != gotValue {
		t.Errorf("Set %d - Got %d", setValue, gotValue)
	}

	setValue = 0
	gotValue = bTile.Get(Height)
	if setValue != gotValue {
		t.Errorf("Set %d - Got %d", setValue, gotValue)
	}
}
