package test

import (
	"legendaryTrain/algorithm"
	"testing"
)

func TestAllSort(t *testing.T) {
	a := []int{3, 2, 5, 9}
	algorithm.PrintPermutations(a, 4, 4)
}
