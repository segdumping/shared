package rand

import (
"math/rand"
	"time"
)

var rGen *rand.Rand

func init() {
	rGen = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GenRand() int {
	return rGen.Int()
}

// [min, max）
func GenRandomRange(min, max int) int {
	if min == max {return min}
	if min == max -1 {return min}
	if max < min {max, min = min, max }

	return rGen.Intn(max - min) + min
}

func GenRandomWeight(values []int) int {
	if len(values) == 0 {
		return -1
	}

	sum := 0
	for index := range  values {
		sum += values[index]
	}

	if sum == 0 {
		return 0
	}

	randValue := rGen.Int() % sum
	min, max := 0, 0
	for index := range values {
		min = max
		value := values[index]
		if value == 0 {
			continue
		}

		max += value
		if randValue >= min && randValue < max {
			return  index
		}
	}

	return  -1
}
