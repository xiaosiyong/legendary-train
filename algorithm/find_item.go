package algorithm

/**
“以时间复杂度O(n)从长度为n的数组中找出所有比左边大比右边的小的元素”。
一开始求出所有的右边最小数组rightMin，然后从左往右判断当前元素是否是左边最大，
如果是则和其相邻的右边最小数（存放于最小数组rightMin）比较，如果小于，则找到了满足条件的元素。
注意：左右两边第一个数不满足条件。
*/
func FindItems(a []int) []int {
	var r []int
	l := len(a)
	t := make([]int, l)
	min := a[l-1]
	for i := l - 1; i >= 0; i-- {
		if a[i] < min {
			min = a[i]
		}
		t[i] = min
	}
	max := a[0]
	for i := 0; i < l-1; i++ {
		if a[i] > max {
			max = a[i]
			if a[i] < t[i+1] {
				r = append(r, a[i])
			}
		}
	}
	return r
}

/***
matrix = [
[1,3,5,7],
[10,11,16,20],
[23,30,34,50]
]
target = 3 输出 true
target = 2 输出 false
*/
func FindItemFromMatrix(a [][]int, s int) bool {
	l := len(a)
	if l < 1 {
		return false
	}
	if l == 1 {

	}

	i, r := findIndex(a, s, 0, len(a))
	if r {
		return true
	}
	if i == -1 {
		return false
	}
	i = binarySearch(a[i], 0, len(a[i]), s)
	return i != -1
}

func findIndex(a [][]int, f, s, e int) (int, bool) {

	m := s + (e-s)>>1
	if a[m][0] == f {
		return m, true
	}
	return 0, false
}

func binarySearch(a []int, s, e, f int) int {
	if s >= e {
		return -1
	}
	if e >= len(a) {
		return -1
	}
	if a[s] > f || a[e] < f {
		return -1
	}
	m := s + (e-s)>>1
	if a[m] == f {
		return m
	} else if a[m] > f {
		return binarySearch(a, s, m, f)
	} else {
		return binarySearch(a, m+1, e, f)
	}
}

/***
给你一个数组，求一个k值，使得前k个数的方差 + 后面n-k个数的方差最小 ，时间复杂度可以到O(n)。
如果不考虑方差的概念，这题可以简化为 “给一个数组，求一个k值，使得前k个数的和 + 后面n-k个数的和最小”。
举例， 如 nums = [1,3,2,4]，我们可以先从左向右求出各个子段和 [1,4,6,10]，然后再从右向左求出各个子段和 [4,6,9,10]，我们发现对应的子段和为 1 -> 9, 4 -> 6, 6 -> 4。因此，我们只需要正反遍历数组两次，就可以求得结果。
时间复杂度：O(n)，空间复杂度 O(n)
*/
func FindIndex(a []int) int {
	var r int
	l := len(a)
	if l > 0 {
		var left, right []int
		var sum, squareSum int
		for i := 0; i < l; i++ {
			sum += a[i]
			squareSum += a[i] * a[i]
			left[i] = (squareSum / (i + 1)) - (sum / (i + 1)) ^ 2
		}
		sum = 0
		squareSum = 0
		for j := l - 1; j > 0; j-- {
			sum += a[j]
			squareSum += a[j] ^ 2
			right[j] = (squareSum / (j - 1)) - (sum / (j - 1)) ^ 2
		}
		r = left[0] + right[0]
		var t int
		for i := 0; i < l-1; i++ {
			if r > left[i]+right[i+1] {
				r = left[i] + right[i+1]
				t = i + 1
			}
		}
		return t

	}
	return r
}
