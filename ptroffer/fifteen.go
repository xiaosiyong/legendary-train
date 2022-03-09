package main

import "fmt"

func main() {
	t := []int{5, 1, 3, 4, 5, 2}
	fmt.Println(repeatedItemNoChange(t, 6))
	fmt.Println()
	fmt.Println(numberOf1(5))
	fmt.Println(numberOf2(5))
	fmt.Println(numberOf2(8))
	fmt.Println(8 & 7)
	fmt.Println(16 & 15)

}

//正整数
func numberOf1(n int) int {
	var r int
	for n != 0 {
		if n&1 > 0 {
			r++
		}
		n = n >> 1
	}
	return r
}

func numberOf2(n int) int {
	var r int
	for n != 0 {
		n = n & (n - 1)
		r++
	}
	return r
}
