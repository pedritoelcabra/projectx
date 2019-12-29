package randomizer

import (
	"math/rand"
	"time"
)

type RandomizerClass struct {
	rand *rand.Rand
}

var Randomizer = RandomizerClass{}

func NewSeed() int {
	return 8730
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(10000)
}

func SetSeed(seed int) {
	source := rand.NewSource(int64(seed))
	Randomizer.rand = rand.New(source)
}

func RandomInt(min, max int) int {
	if min > max {
		return RandomInt(max, min)
	}
	if min < 0 {
		offset := -min
		offsetResult := RandomInt(min+offset, max+offset)
		return min + offsetResult
	}
	randomInt := Randomizer.rand.Intn(max)
	return min + randomInt
}

func RandomFloat() float64 {
	return Randomizer.rand.Float64()
}

func PercentageRoll(percentage int) bool {
	roll := RandomInt(0, 100)
	return roll <= percentage
}
