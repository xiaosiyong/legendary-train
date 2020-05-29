package test

import (
	"fmt"
	"legendaryTrain/algorithm"
	"testing"
)

func TestFindSqr(t *testing.T) {
	//fmt.Println(algorithm.FindSqr(13,0.000001))
	//fmt.Println(3.605551272630*3.605551272630)
	a := []int{2, 8, 10}
	fmt.Println(algorithm.FindItems(a))
	fmt.Println(algorithm.FindItems(a))

}

func TestFindItem(t *testing.T) {
	a := [][]int{
		{1, 1, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 1, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
	}
	//fmt.Println(algorithm.GetIslandQuantity(a))
	fmt.Println(algorithm.GetContinuousQuantity(a))
}
