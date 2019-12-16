package test

import (
	"fmt"
	"legendaryTrain/datastruct"
	"testing"
)

func TestArrayStack(t *testing.T) {
	a := datastruct.InitArrayStack(5)
	fmt.Println(a.Length)
	a.Push("Hello")
	a.Push("it's me")
	fmt.Println(a.Length, a.Count)
	fmt.Println(a.Pop())

}
