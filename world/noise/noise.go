package noise

import "github.com/aquilax/go-perlin"

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
	noise := n.p.Noise2D(float64(x)/10, float64(y)/10)
	adjustedNoise := int(noise * 1000)
	return adjustedNoise
}
