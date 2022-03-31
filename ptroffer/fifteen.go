package main

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
