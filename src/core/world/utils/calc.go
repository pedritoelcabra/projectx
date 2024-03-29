package utils

import "fmt"

func AbsInt(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func NumberFormat(f float64) string {
	return fmt.Sprintf("%.2f", f)
}
