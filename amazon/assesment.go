package amazon

import "fmt"

/*
1、[1,2,3,4,5,6]分成两部分，一部分含有最少的元素，其和大于剩下一部分的和
2、1->2->3->4   从首尾往中间遍历，找和最大的，要求空间复杂度O(1)
3、计算单词的不重复值，比如good，从一个字母，到两个，到整个单词
*/

func DivideVolume(a []int) []int {
	var sum int
	for _, v := range a {
		sum += v
	}
	fmt.Println(a)
	bucketSort(a)
	fmt.Println(a)
	var temp int
	var setA []int
	for i := len(a) - 1; i >= 0; i-- {
		if temp < sum-temp {
			setA = append(setA, a[i])
			temp += a[i]
		} else {
			return setA
		}
	}
	return setA
}

func bucketSort(a []int) {
	//max := 100
	bucket := [101]int{}
	for _, v := range a {
		bucket[v]++
	}
	index := 0
	for i := 0; i < 101; i++ {
		for j := 1; j <= bucket[i]; j++ {
			a[index] = i
			index++
		}
	}

}

func countSort(a []int32) {

}
