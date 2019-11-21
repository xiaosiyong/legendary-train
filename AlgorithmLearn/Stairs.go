package AlgorithmLearn

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
	prepre := 1
	pre := 2
	for i := 3; i <= n; i++ {
		ret = pre + prepre
		prepre = pre
		pre = ret
	}
	return ret
}
