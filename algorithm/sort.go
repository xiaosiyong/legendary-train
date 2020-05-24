package algorithm

//插入排序
func InsertSort(a []int, n int) []int {
	if n <= 1 {
		return a
	}
	for i := 1; i < n; i++ {
		v := a[i]
		// 查找插入的位置
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > v {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = v
	}
	return a
}

//冒泡排序
func BubbleSort(a []int, l int) []int {
	if l <= 1 {
		return a
	}
	for i := 0; i < l; i++ {
		flag := false
		for j := 0; j < l-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return a
}

//选择排序
func SelectSort(a []int, l int) []int {
	for i := 0; i < l; i++ {
		temp := i
		for j := i + 1; j < l; j++ {
			if a[j] < a[temp] {
				temp = j
			}
		}
		a[i], a[temp] = a[temp], a[i]

	}
	return a
}

//归并排序
func MergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	m := len(a) / 2
	l := MergeSort(a[:m])
	r := MergeSort(a[m:])
	return merge(l, r)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

//快速排序
func QuickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	i := partition(a, lo, hi)
	QuickSort(a, lo, i-1)
	QuickSort(a, i+1, hi)

}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[j], a[i] = a[i], a[j]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

//桶排序（基数） 必须是有界数据
func BucketSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	t := [100]int{}
	for i := 0; i < len(a); i++ {
		t[a[i]]++
	}
	index := 0
	for i := 0; i < 100; i++ {
		for j := 1; j <= t[i]; j++ {
			a[index] = i
			index++
		}
	}
	return a
}

//计数排序
func CountSort(a []int) []int {
	max := 0
	l := len(a)
	min := 0
	for i := 0; i < l; i++ {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
	}

	t := make([]int, max-min+1)
	for i := 0; i < l; i++ {
		t[a[i]] += 1
	}
	j := 0
	for i := 0; i < max-min+1; i++ {
		for t[i] > 0 {
			a[j] = i
			j++
			t[i]--
		}
	}
	return a
}

func QuickSort2(a []int) []int {
	if len(a) < 2 {
		return a
	} else {
		pivot := a[0]
		var less, greater []int
		for _, v := range a[1:] {
			if v > pivot {
				greater = append(greater, v)
			} else {
				less = append(less, v)
			}
		}
		var result []int
		result = append(result, QuickSort2(less)...)
		result = append(result, pivot)
		result = append(result, QuickSort2(greater)...)
		return result
	}

}

func TopN(a []int, l, r, n int) int {
	if l == r {
		return a[r]
	}
	p := partitionMax(a, l, r)
	if p-l+1 > n {
		return TopN(a, l, p-1, n)
	} else if p-l+1 == n {
		return a[p]
	} else {
		return TopN(a, p+1, r, n-p+l-1)

	}

}

func partitionMax(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] > pivot {
			a[j], a[i] = a[i], a[j]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

type Graph struct {
}
