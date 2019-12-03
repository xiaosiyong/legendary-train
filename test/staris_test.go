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
