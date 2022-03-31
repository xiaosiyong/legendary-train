package main

import "fmt"

func main() {
	//t := []int{5, 1, 3, 4, 5, 2}
	//fmt.Println(repeatedItemNoChange(t, 6))
	//fmt.Println()
	//fmt.Println(numberOf1(5))
	//fmt.Println(numberOf2(5))
	//fmt.Println(numberOf2(8))
	//fmt.Println(8 & 7)
	//fmt.Println(16 & 15)

	a := [][]int{
		{1, 2, 3}, {4, 5, 6}, {8, 10, 12},
	}
	fmt.Println(existInArray(a, 9))
}
