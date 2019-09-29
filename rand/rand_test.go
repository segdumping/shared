package rand

import "testing"

func TestGenRand(t *testing.T) {
	t.Logf("gen rand: %d", GenRand())
}

func TestGenRandomRange(t *testing.T) {
	t.Logf("[1 2] -> %d", GenRandomRange(1, 2))
	t.Logf("[2 2] -> %d", GenRandomRange(2, 2))
	t.Logf("[1 5] -> %d", GenRandomRange(1, 5))
	t.Logf("[4 2] -> %d", GenRandomRange(4, 2))
}

func TestGenRandomWeight(t *testing.T) {
	var weight = []int {
		10,
		30,
		50,
		5,
		5,
	}

	t.Logf("gen random weight: %d", GenRandomWeight(weight))
}