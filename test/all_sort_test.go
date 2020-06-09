package test

import (
	"legendaryTrain/algorithm"
	"testing"
)

func TestAllSort(t *testing.T) {
	a := []int{1, 2, 3, 4}
	algorithm.PrintPermutations(a, 4, 4)
}
