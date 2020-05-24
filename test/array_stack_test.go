package test

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	//a := datastruct.InitArrayStack(5)
	//fmt.Println(a.Length)
	//a.Push("Hello")
	//a.Push("it's me")
	//fmt.Println(a.Length, a.Count)
	//fmt.Println(a.Pop())
	var a [][]int
	var t1 []int
	t1 = append(t1, 1)
	t1 = append(t1, 2)
	t1 = append(t1, 3)
	a = append(a, t1)

	var t2 []int
	t2 = append(t2, 5)
	t2 = append(t2, 6)
	a = append(a, t2)

	fmt.Println(len(a[0]))

}
