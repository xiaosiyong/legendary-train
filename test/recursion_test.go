package test

import (
	"fmt"
	"legendaryTrain/algorithm"
	"testing"
)

func TestCalculateTotal1(t *testing.T) {
	fmt.Println("队列测试结束")
	fmt.Println("总共10个台阶，共有走法：", algorithm.CalculateTotal1(10))
}

func TestCalculateTotal2(t *testing.T) {
	fmt.Println("总共10个台阶，共有走法：", algorithm.CalculateTotal2(10))
}

func TestCalculateTotal3(t *testing.T) {
	fmt.Println("总共10个台阶，共有走法：", algorithm.CalculateTotal3(10))
}

func TestCalculateFactorial(t *testing.T){
	fmt.Println("N的阶乘，N=",10,",值：",algorithm.CalculateFactorial(10))
}

func TestRangeAll(t *testing.T){
	//r := algorithm.NewRangeType(3)
	//for i:=0;i<3;i++{
	//	r.Value[i] = i+1
	//}
	//r.RangAll(0)

	slice := algorithm.NewRangeType(3)
	slice.Value[0] = "a"
	slice.Value[1] = "b"
	slice.Value[2] = "c"
	//slice.Value[3] = "a"

	slice.RangAll(0)
}