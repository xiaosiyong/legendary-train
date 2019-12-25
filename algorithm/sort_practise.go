package algorithm

func bubleSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	var flag bool
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return a
}

func insertSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}
	for i := 1; i < l; i++ {
		v := a[i]  //把值存起来 位置空出来
		j := i - 1 //向左去遍历
		for ; j >= 0; j-- {
			if a[j] > v {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = a[i]
	}
	return a
}

func selectSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}
	for i := 0; i < l; i++ {
		t := i
		for j := i + 1; j < l; j++ {
			if a[j] < a[t] {
				t = j
			}
		}
		a[i], a[t] = a[t], a[i]
	}
	return a
}
