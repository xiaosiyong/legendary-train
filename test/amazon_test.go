package test

import (
	"fmt"
	"legendaryTrain/amazon"
	"testing"
)

func TestDivideVolume(t *testing.T) {
	a := []int{15, 2, 8, 4, 3, 7, 5}
	result := amazon.DivideVolume(a)
	fmt.Println(result)
}
