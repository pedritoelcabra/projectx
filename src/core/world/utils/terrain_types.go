package utils

type TerrainTypes int

const (
	MountainHeight = 300
	HillHeight     = 150
)

const (
	ScarceForestBioMass = 200
	ForestBioMass       = 350
)

const (
	BasicHills     = 1
	BasicGrass     = 2
	BasicWater     = 3
	BasicMountain  = 4
	BasicDeepWater = 5

	TundraHills     = 11
	TundraGrass     = 12
	TundraWater     = 13
	TundraMountain  = 14
	TundraDeepWater = 15

	DesertHills     = 21
	DesertGrass     = 22
	DesertWater     = 23
	DesertMountain  = 24
	DesertDeepWater = 25

	BiomeTemperate = 100
	BiomeTundra    = 101
	BiomeDesert    = 102

	BorderN  = 1000
	BorderNE = 1001
	BorderSE = 1002
	BorderS  = 1003
	BorderSW = 1004
	BorderNW = 1005

	GreenOverlay = 1010
	RedOverlay   = 1011
)

var movementCosts = make(map[int]float64)

func DirectionToBorder(dir int) int {
	return dir + 1000
}

func MovementCost(terrain int) float64 {
	switch terrain {
	case BasicHills:
		return 1.8
	case BasicDeepWater:
		return 1000.0
	case BasicMountain:
		return 1000.0
	case BasicGrass:
		return 1.0
	case BasicWater:
		return 2.0
	case TundraHills:
		return 3.0
	case TundraDeepWater:
		return 1000.0
	case TundraMountain:
		return 1000.0
	case TundraGrass:
		return 1.5
	case TundraWater:
		return 1.5
	case DesertHills:
		return 2.0
	case DesertDeepWater:
		return 1000.0
	case DesertMountain:
		return 1000.0
	case DesertGrass:
		return 1.0
	case DesertWater:
		return 2.0
	}
	return 0.0
}
