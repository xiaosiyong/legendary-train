package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"legendaryTrain/algorithm"
	"math/rand"
	"time"
)

type locDesc struct {
	Key    string   `json:"key"`
	Params []string `json:"params,omitempty"`
}

func GetLocDesc(key string, params ...string) string {
	t, _ := json.Marshal(&locDesc{
		Key:    key,
		Params: params,
	})

	return string(fmt.Sprintf("`%s`", t))

}

func main() {

	var testMap map[int64]string
	r, ok := testMap[2]
	fmt.Println(r, ok)

	fmt.Println(time.Now().Add(10 * time.Second).Unix())

	err := errors.New("error hahh")
	fmt.Println(fmt.Sprintf("error:%s", err))

	algorithm.FindAllSuShu(100)

	//go func() {
	//
	//	for {
	//
	//		fmt.Println("1 seconds past")
	//		time.Sleep(20 * time.Millisecond)
	//
	//	}
	//
	//}()
	//t := time.NewTicker(2 * time.Second)
	//for {
	//	select {
	//	case <-t.C:
	//		fmt.Println("2 seconds past")
	//	}
	//}

}

func FindKthLargest(array []int, k int) int {
	length := len(array)
	if length == 0 {
		return 0
	}
	//如果只有一个元素 就返回原
	if length == 1 {
		return array[0]
	}
	return quickSelect(array, 0, length-1, k)
}

func partitionArray(array []int, start, end, partitionIndex int) int {
	currentValue := array[partitionIndex]
	array[partitionIndex], array[start] = array[start], array[partitionIndex]
	partition := end
	for i := end; i > 0; i-- {
		if array[i] < currentValue {
			array[i], array[partition] = array[partition], array[i]
			partition--
		}
	}
	array[partition], array[start] = array[start], array[partition]
	return partition

}

func quickSelect(array []int, start, end, k int) int {
	//如果区间内只有一个元素直接返回
	if start == end {
		return array[start]
	}
	//取随机元素
	randIndex := start + rand.Intn(end-start)
	position := partitionArray(array, start, end, randIndex)
	if position == k-1 {
		return array[position]
	} else if position > k-1 {
		return quickSelect(array, 0, position-1, k)
	} else {
		return quickSelect(array, position+1, end, k)
	}

}
