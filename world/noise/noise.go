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
