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
