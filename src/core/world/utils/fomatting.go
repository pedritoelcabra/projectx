package utils

import (
	"fmt"
	"strconv"
)

func FormatCount(count int) string {
	if count < 1000 {
		return strconv.Itoa(count)
	}
	countFloat := float64(count)
	if count < 1000000 {
		return fmt.Sprintf("%.1f", countFloat/1000) + "K"
	}
	return fmt.Sprintf("%.1f", countFloat/1000000) + "K"
}
