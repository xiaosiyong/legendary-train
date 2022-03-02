package main

//找出数组内重复的元素，可以更改原数组，长度为n，在1~n-1范围内
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

//找出数组内重复的元素，不能更改原数组，长度为n，在1~n-1范围内
func repeatedItemNoChange(a []int, n int) int {
	return -1
}
