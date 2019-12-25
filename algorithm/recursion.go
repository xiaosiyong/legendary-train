package algorithm

import "fmt"

/*
N 个台阶 每次可以走一步 也可以走两步 问总共有多少种走法

*/
var (
	m = make(map[int]int)
)

func CalculateTotal1(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return CalculateTotal1(n-1) + CalculateTotal1(n-2)
}

func CalculateTotal2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if va, ok := m[n]; ok {
		return va
	}
	ret := CalculateTotal2(n-1) + CalculateTotal2(n-2)
	m[n] = ret
	return ret

}

func CalculateTotal3(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	ret := 0
	prepare := 1
	pre := 2
	for i := 3; i <= n; i++ {
		ret = pre + prepare
		prepare = pre
		pre = ret
	}
	return ret
}

func CalculateFactorial(n int) int {
	if n <= 2 {
		return n
	}
	if va, ok := m[n]; ok {
		return va
	}
	result := n * CalculateFactorial(n-1)
	m[n] = result
	return result
}

type RangeType struct {
	Value []interface{}
}

func NewRangeType(n int) *RangeType {
	return &RangeType{
		make([]interface{}, n),
	}
}

//TODO 复习最大子序列
func (r *RangeType) RangAll(start int) {
	len := len(r.Value)
	if start == len-1 {
		fmt.Println(r.Value)
	}
	for i := start; i < len; i++ {
		if i == start || r.Value[i] != r.Value[start] { //重复元素则不排序
			r.Value[i], r.Value[start] = r.Value[start], r.Value[i]
			r.RangAll(start + 1)
			r.Value[i], r.Value[start] = r.Value[start], r.Value[i]
		}
	}
}
