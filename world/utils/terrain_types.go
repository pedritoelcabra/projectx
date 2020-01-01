package utils

type TerrainTypes int

const (
	BasicHills     = 1
	BasicGrass     = 2
	BasicWater     = 3
	BasicMountain  = 4
	BasicDeepWater = 5
	BorderN        = 1000
	BorderNE       = 1001
	BorderSE       = 1002
	BorderS        = 1003
	BorderSW       = 1004
	BorderNW       = 1005
)

var movementCosts = make(map[int]float64)

func DirectionToBorder(dir int) int {
	return dir + 1000
}

func MovementCost(terrain int) float64 {
	switch terrain {
	case BasicHills:
		return 2.0
	case BasicDeepWater:
		return 1000.0
	case BasicMountain:
		return 1000.0
	case BasicGrass:
		return 1.0
	case BasicWater:
		return 2.0
	}
	return 0.0
}
