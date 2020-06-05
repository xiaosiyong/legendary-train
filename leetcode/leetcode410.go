package leetcode

import "math"

/***
给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。
*/
func SplitArray(nums []int, m int) int {
	l := len(nums)
	if l > 0 {
		//init the sum array
		s := make([]int, l+1)
		for i := 0; i < l; i++ {
			s[i+1] = s[i] + nums[i]
		}
		//状态定义：f[i][j]表示nums[0] ~ nums[j]共j+1个元素划分为i组的和的最大最小值
		f := make([][]int, l+1)
		for i := 0; i <= l; i++ {
			f[i] = make([]int, m+1)
			for j := 0; j <= m; j++ {
				f[i][j] = math.MaxInt32
			}
		}
		f[0][0] = 0
		for i := 1; i <= l; i++ {
			for j := 1; j <= m; j++ {
				for t := 0; t < i; t++ {
					f[i][j] = minInt(f[i][j], maxInt(s[i]-s[t], f[t][j-1]))
				}
			}
		}
		return f[l][m]
	}
	return 0
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func SplitArray2(nums []int, m int) int {
	var l, r int
	for i := 0; i < len(nums); i++ {
		r += nums[i]
		if l < nums[i] {
			l = nums[i]
		}
	}
	ans := r
	for l <= r {
		mid := l + (r-l)>>1
		var s int
		c := 1
		for i := 0; i < len(nums); i++ {
			if s+nums[i] > mid {
				c++
				s = nums[i]
			} else {
				s += nums[i]
			}
		}
		if c <= m {
			ans = minInt(ans, mid)
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}
