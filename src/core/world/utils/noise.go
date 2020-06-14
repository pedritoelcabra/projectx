package utils

import (
	"github.com/aquilax/go-perlin"
	"math"
)

const (
	alpha = 2.
	beta  = 2.
	n     = 3
)

type NoiseGenerator struct {
	p *perlin.Perlin
}

var Generator = NoiseGenerator{}

func Seed(seed int) {
	Generator.p = perlin.NewPerlin(alpha, beta, n, int64(seed))
}

func (n *NoiseGenerator) GetHeight(x, y int) int {
	height := 0.0
	ruggedIntensity := math.Abs(n.ApplyFilter(x, y, 100.0, 2000.0) / 1000)
	if ruggedIntensity >= 1.0 {
		ruggedIntensity = 1.0
	}
	rugged := n.ApplyFilter(x+100, y+100, 10.0, 400.0)
	rugged *= ruggedIntensity
	height += rugged
	plaque := n.ApplyFilter(x+10000, y+10000, 100.0, 600.0)
	height += plaque
	baseline := 100.0
	height += baseline
	if height >= 1000 {
		height = 999.0
	}
	return int(height)
}

func (n *NoiseGenerator) GetBiome(x, y int) int {
	val := n.ApplyFilter(x, y-10000, 300.0, 1000.0)
	return int(val)
}

func (n *NoiseGenerator) GetBiomass(x, y int) int {
	val := n.ApplyFilter(x, y+10000, 50.0, 1000.0)
	return int(val)
}

func (n *NoiseGenerator) GetStone(x, y int) int {
	val := n.ApplyFilter(x, y+20000, 50.0, 1000.0)
	return int(val)
}

func (n *NoiseGenerator) GetIron(x, y int) int {
	val := n.ApplyFilter(x, y+30000, 50.0, 1000.0)
	return int(val)
}

func (n *NoiseGenerator) GetCoal(x, y int) int {
	val := n.ApplyFilter(x, y+40000, 50.0, 1000.0)
	return int(val)
}

func (n *NoiseGenerator) ApplyFilter(x, y int, scale float64, intensity float64) float64 {
	noise := n.p.Noise2D(float64(x)/scale, float64(y)/scale)
	adjustedNoise := noise * intensity
	return adjustedNoise
}
