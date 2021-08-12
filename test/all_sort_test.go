package test

import (
	"legendaryTrain/algorithm"
	"testing"
)

func TestAllSort(t *testing.T) {
	a := []int{3, 2, 5,4}
	algorithm.PrintPermutations(a, 4, 4)

	b := []int{4,6,8,7}
	algorithm.AllSort(b,4,4)
}
