package test

import (
	"fmt"
	"legendaryTrain/algorithm"
	"testing"
)

func TestRandomMoney(t *testing.T){
	p := &algorithm.RedPacket{
		TotalQuantity:  10,
		TotalMoney:     100.00,
		RemainQuantity: 10,
		RemainMoney:    100.00,
	}
	var r,temp float64
	for i :=0 ;i <10;i++{
		r = algorithm.RandomMoney(p)
		fmt.Println(r)
		temp += r
	}
	fmt.Printf("总金额%.2f",temp)
}
