package algorithm

import "fmt"

//判断一个数能否由整数个7、13、29构成
func IsDivide(n int) bool {
	items := []int{7, 13, 29}
	if n < 29 {
		if n == 7 || n == 13 || n == 20 || n == 26 || n == 27 {
			return true
		}
	}

	for _, t := range items {
		if n%t == 0 {
			return true
		}
	}

	var a []int
	for i := 0; i <= n; i++ {
		a = append(a, 0)
	}

	for _, t := range items {
		a[t] = 1
	}
	for i := 7; i <= n; i++ {
		for _, t := range items {
			if i-t >= 0 && a[i-t] == 1 {
				a[i] = 1
			}
		}
	}
	fmt.Println(a)
	return a[n] == 1
}
