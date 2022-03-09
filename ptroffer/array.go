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

//找出数组内重复的元素，不能更改原数组，长度为n，在1~n-1范围内
//hint，由于元素的范围在1~n，可以折半查找，看1~n/2内个数，将数组分成两部分，看看数字出现的次数
func repeatedItemNoChange(a []int, n int) int {

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
