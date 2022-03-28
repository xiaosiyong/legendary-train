package test

import (
	"fmt"
	"legendaryTrain/algorithm"
	"testing"
)

func TestIsDivide(t *testing.T) {
	fmt.Println(algorithm.IsDivide(100))
	a := []int{10, 100, 3, 4, 5, 5, 6, 8, 11, 60, 100}
	bucketSort(a)
	fmt.Println(a)
}

//最大是100
func bucketSort(a []int) {
	temp := [101]int{}
	for _, v := range a {
		temp[v]++
	}
	var index int
	for i := 0; i < 101; i++ {
		for j := 1; j <= temp[i]; j++ {
			a[index] = i
			index++
		}
	}

}
