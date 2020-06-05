package test

import (
	"fmt"
	"legendaryTrain/leetcode"
	"testing"
)

func TestLeetCode15(t *testing.T) {
	a := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	result := leetcode.ThreeSum2(a)
	fmt.Println(result)

	fmt.Println(5 >> 1)
	//slice 与 array 复习
	var temp []int
	fmt.Println("var temp []int,temp 是 nil吗？，", temp == nil)
	temp = append(temp, 1)
	fmt.Println(temp, len(temp), cap(temp))

	m := make([]int, 3)
	fmt.Println(m, len(m), cap(m))
	m = append(m, 1)
	fmt.Println(m, len(m), cap(m))
	m = append(m, 12)
	fmt.Println(m, len(m), cap(m))

	var array [3]int
	fmt.Println(array[1])
	array[0] = 1
	fmt.Println(array)

}

func TestFindKthLargest(t *testing.T) {
	b := []int{-1, 2, 0}
	fmt.Println(leetcode.FindKthLargest(b, 2))
	fmt.Println(b)
}

func Test(t *testing.T) {
	a := []int{4, 5, 6, 7, 8, 1, 2, 3}
	fmt.Println(leetcode.Search(a, 6))
}

func TestBfsPrint(t *testing.T) {
	t4 := &leetcode.TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	t5 := &leetcode.TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	t6 := &leetcode.TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	t7 := &leetcode.TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	t2 := &leetcode.TreeNode{
		Val:   2,
		Left:  t6,
		Right: t7,
	}
	t3 := &leetcode.TreeNode{
		Val:   3,
		Left:  t4,
		Right: t5,
	}
	t1 := &leetcode.TreeNode{
		Val:   1,
		Left:  t2,
		Right: t3,
	}
	fmt.Println(leetcode.PrintByLevel(t1))
}

func TestLeetcode01(t *testing.T) {
	m := []int{3, 2, 4}
	fmt.Println(leetcode.TwoSum2(m, 6))
}

func TestRotateSlice(t *testing.T) {
	m := [][]int{{3, 2, 4}, {1, 3, 4}, {1, 2}}
	fmt.Println(m[0][1])
	fmt.Println(len(m[0]))
}

func TestLeetCode84(t *testing.T) {
	a := []int{2, 1, 5, 6, 2, 3}
	//fmt.Println(leetcode.LargestRectangleArea1(a))
	//fmt.Println(leetcode.LargestRectangleArea2(a))
	//fmt.Println(leetcode.LargestRectangleArea3(a))
	fmt.Println(leetcode.LargestRectangleArea4(a))
}

func TestLeetCode85(t *testing.T) {
	a := [][]int{
		{1, 0, 0, 0, 1},
		{1, 1, 0, 0, 1},
		{1, 1, 1, 0, 1},
		{1, 1, 0, 1, 1},
		{1, 0, 1, 0, 1},
	}
	//fmt.Println(leetcode.MaximalRectangle1(a))
	//fmt.Println(leetcode.MaximalRectangle2(a))
	fmt.Println(leetcode.MaximalRectangle3(a))
}

func TestLeetCode410(t *testing.T) {
	a := []int{7, 2, 5, 10, 8}
	fmt.Println(leetcode.SplitArray2(a, 2))
}
