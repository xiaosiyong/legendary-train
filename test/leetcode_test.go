package test

import (
	"fmt"
	"legendaryTrain/leetcode"
	"testing"
)

func TestLeetCode15(t *testing.T){
	a := []int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}
	result :=leetcode.ThreeSum2(a)
	fmt.Println(result)

	fmt.Println(5>>1)
	//slice 与 array 复习
	var temp []int
	fmt.Println("var temp []int,temp 是 nil吗？，",temp==nil)
	temp = append(temp,1)
	fmt.Println(temp,len(temp),cap(temp))

	m := make([]int,3)
	fmt.Println(m,len(m),cap(m))
	m = append(m,1)
	fmt.Println(m,len(m),cap(m))
	m = append(m,12)
	fmt.Println(m,len(m),cap(m))

	var array [3]int
	fmt.Println(array[1])
	array[0] =1
	fmt.Println(array)


}

func TestFindKthLargest(t *testing.T){
	b := []int{-1,2,0}
	fmt.Println(leetcode.FindKthLargest(b,2))
	fmt.Println(b)
}


func Test(t *testing.T){
	a := []int{4,5,6,7,8,1,2,3}
	fmt.Println(leetcode.Search(a,6))

}


