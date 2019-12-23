package movement

import "math"

func AdvanceAlongLine(x1, y1, x2, y2, maxDis float64) (x3, y3 float64) {
	dis := CalculateDistance(x1, y1, x2, y2)
	if dis <= maxDis {
		return x2, y2
	}
	ratio := maxDis / dis
	x3 = x1 + (ratio * (x2 - x1))
	y3 = y1 + (ratio * (y2 - y1))
	return
}

func CalculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
