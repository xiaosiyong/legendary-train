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
	fmt.Println(algorithm.FindItem(a))
	fmt.Println(algorithm.FindItems(a))

}
