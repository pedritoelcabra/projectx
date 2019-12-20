package noise

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

func New(seed int) *NoiseGenerator {
	aGenerator := &NoiseGenerator{}
	aGenerator.p = perlin.NewPerlin(alpha, beta, n, int64(seed))
	return aGenerator
}

func (n *NoiseGenerator) GetHeight(x, y int) int {
	height := 0.0
	rugged_intensity := math.Abs(n.ApplyFilter(x, y, 100.0, 1000.0) / 1000)
	rugged := n.ApplyFilter(x+100, y+100, 10.0, 200.0)
	rugged *= rugged_intensity
	height += rugged
	plaque := n.ApplyFilter(x+10000, y+10000, 50.0, 700.0)
	height += plaque
	return int(height)
}

func (n *NoiseGenerator) ApplyFilter(x, y int, scale float64, intensity float64) float64 {
	noise := n.p.Noise2D(float64(x)/scale, float64(y)/scale)
	adjustedNoise := noise * intensity
	return adjustedNoise
}
