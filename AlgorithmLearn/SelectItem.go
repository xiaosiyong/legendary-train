package AlgorithmLearn

//1,2,3,4,5,5,5,7,8,9
func SelectFirstIndexOfItem(a []int, n, b int) int {
	if b > a[n-1] {
		return -1
	}

	l := 0
	h := n - 1
loop:
	m := l + ((h - l) >> 1)
	if a[m] >= b {
		h = m - 1
	} else {
		l = m + 1
	}
	if l <= h {
		goto loop

	}
	if l < n && a[l] == b {
		return l
	}
	return -1
}

func SelectFirstIndexOfItem2(a []int, n, b int) int {
	if b > a[n-1] {
		return -1
	}

	l := 0
	h := n - 1
loop:
	m := l + ((h - l) >> 1)
	if a[m] > b {
		h = m - 1
	} else if a[m] < b {
		l = m + 1
	} else {
		if m == 0 || a[m-1] != b {
			return m
		}
	}
	goto loop
	return -1
}

//二分查找
func BinarySearch(a []int, l, h, b int) int {
	if len(a) < 1 {
		return -1
	}
	if a[len(a)-1] < b || a[0] > b {
		return -1
	}
	if l >= h {
		return -1
	}

	m := l + ((h - l) >> 1)
	t := a[m]
	if t == b {
		return m
	}

	if t > b {
		return BinarySearch(a, l, m-1, b)
	} else {
		return BinarySearch(a, m+1, h, b)
	}
}
