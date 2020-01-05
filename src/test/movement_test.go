package test

import (
	"github.com/pedritoelcabra/projectx/src/world/utils"
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	distances := []struct {
		x1, y1, x2, y2, dist float64
	}{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, -1, 1, 1.41},
		{-10, -10, 10, 10, 28.28},
	}
	for _, c := range distances {
		measured := math.Round(utils.CalculateDistance(c.x1, c.y1, c.x2, c.y2)*100) / 100
		if measured != c.dist {
			t.Errorf("%.2f / %.2f to %.2f / %.2f measured %.2f, want %.2f", c.x1, c.y1, c.x2, c.y2, measured, c.dist)
		}
	}
}

func TestMovementAlongLine(t *testing.T) {
	distances := []struct {
		x1, y1, x2, y2, speed, x3, y3 float64
	}{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 5, 5, 100, 5, 5},
		{0, 0, 0, 100, 100, 0, 100},
		{0, 0, 0, 100, 50, 0, 50},
		{0, 0, 0, -100, 50, 0, -50},
		{0, 0, 100, 100, 50, 35.36, 35.36},
		{0, 0, 100, 200, 100, 44.72, 89.44},
		{100, 0, 100, 200, 100, 100, 100},
	}
	for _, c := range distances {
		x3, y3 := utils.AdvanceAlongLine(c.x1, c.y1, c.x2, c.y2, c.speed)
		x3 = math.Round(x3*100) / 100
		y3 = math.Round(y3*100) / 100
		if x3 != c.x3 || y3 != c.y3 {
			t.Errorf("result %.2f / %.2f want %.2f / %.2f", x3, y3, c.x3, c.y3)
		}
	}
}
