package main

//找出数组内重复的元素，数组内元素范围在1~n-1范围内，可以更改原数组，长度为n，
func repeatedItem(a []int, n int) int {
	if len(a) <= 1 || n <= 1 {
		return -1
	}
	for i := 0; i < n; {
		if a[i] != i {
			if a[i] == a[a[i]] {
				return a[i]
			}
			a[i], a[a[i]] = a[a[i]], a[i]
		} else {
			i++
		}
	}
	return -1
}

/*
hint：
数组元素的取值范围是 [1, n]，对该范围对半划分，分成 [1, middle], [middle+1, n]。
计算数组中有多少个(count)元素落在 [1, middle] 区间内，如果 count 大于 middle-1+1，那么说明这个范围内有重复元素，否则在另一个范围内。
继续对这个范围对半划分，继续统计区间内元素数量。
时间复杂度 O(n * log n)，空间复杂度 O(1)。
*/
func repeatedItemNoChange(a []int, n int) int {
	if len(a) == 0 {
		return -1
	}
	start := 1
	end := n - 1
	for end >= start {
		middle := (end-start)>>1 + start
		count := findItemCount(a, start, middle)
		if end == start {
			if count > 1 {
				return start
			}
			break
		} else {
			//说明出现的在小的的那一半里
			if count > middle-start+1 {
				end = middle
			} else {
				start = middle + 1
			}
		}
	}
	return -1
}

//找出子数组中出现的次数
func findItemCount(a []int, min, max int) int {
	if len(a) == 0 {
		return 0
	}
	var c int
	for _, v := range a {
		if v >= min && v <= max {
			c++
		}
	}
	return c
}

//二位数组中出现的元素
func existInArray(a [][]int, target int) bool {
	row := len(a[0])
	return false
}
