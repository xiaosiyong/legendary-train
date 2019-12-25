package test

import (
	"fmt"
	"legendaryTrain/algorithm"
	"testing"
)

func TestSort(t *testing.T) {
	//a := []int{4, 5, 6, 3, 2, 1}
	//fmt.Println("冒泡排序之前：", a)
	//sorted := algorithm.BubbleSort(a, 6)
	//fmt.Println("冒泡排序之后：", sorted)
	//fmt.Println("")
	//
	//b := []int{4, 5, 6, 3, 2, 1}
	//fmt.Println("插入排序之前：", b)
	//fmt.Println("插入排序之后：", algorithm.InsertSort(b, 6))
	//fmt.Println("")
	//
	//c := []int{4, 5, 6, 3, 2, 1}
	//fmt.Println("选择排序之前：", c)
	//fmt.Println("选择排序之后：", algorithm.SelectSort(c, 6))
	//fmt.Println("")
	//
	//d := []int{4, 5, 6, 1, 3, 2, 7}
	//fmt.Println("归并排序之前：", d)
	//fmt.Println("归并排序之后：", algorithm.MergeSort(d))
	//fmt.Println("")
	//
	//e := []int{4, 5, 6, 1, 3, 2, 7}
	//fmt.Println("桶排序之前：", e)
	//fmt.Println("桶排序之后：", algorithm.BucketSort(e))
	//fmt.Println("")

	//f := []int{6, 4, 5, 7, 2, 1, 3}
	//fmt.Println("快速排序1之前：", f)
	//algorithm.QuickSort(f, 0, 6)
	//fmt.Println("快速排序1之后：", f)
	//fmt.Println("")
	//
	//g := []int{4, 5, 6, 1, 3, 2, 7}
	//fmt.Println("快速排序2之前：", g)
	//fmt.Println("快速排序2之后：", algorithm.QuickSort2(g))
	//fmt.Println("")
	//
	//h := []int{4, 5, 6, 1, 3, 2, 7}
	//fmt.Println(algorithm.TopN(h, 0, 6, 2))
	//
	//temp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100, 102, 201, 400}
	//fmt.Println("100的位置是：", algorithm.BinarySearch(temp, 0, 12, 6))
	//
	//a1 := []int{1, 2, 3, 4, 5, 5, 5, 7, 8, 9}
	//fmt.Println("第一个5的位置：", algorithm.SelectFirstIndexOfItem(a1, len(a1), 5))

	t3 := []int{4, 5, 3, 1, 9, 6, 8}
	fmt.Println(algorithm.CountSort(t3))
}
