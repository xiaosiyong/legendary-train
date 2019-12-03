package test

import (
	"fmt"
	"legendaryTrain/algorithm"
	"testing"
)

func TestGetMaxSum(t *testing.T) {
	t1 := []int{1, 2, 3, 4, 5, -100, 101, 2, 12, -20}
	fmt.Println(algorithm.GetMaxSum(t1))
	fmt.Println(algorithm.GetMaxSum2(t1))

	t2 := []int{-11, -2}
	fmt.Println(algorithm.GetMaxSum(t2))
	fmt.Println(algorithm.GetMaxSum2(t2))
}
