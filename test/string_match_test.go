package test

import (
	"fmt"
	"legendaryTrain/algorithm"
	"testing"
)

func TestStringMatch(t *testing.T) {
	mainStr := "abcdabaabcdefacb"
	searchStr1 := "bddeg"
	searchStr2 := "abc"
	searchStr3 := "acb"
	searchStr4 := "aabcdf"
	fmt.Println(searchStr1, "在", mainStr, "匹配的起始位置：", algorithm.BM([]byte(mainStr), []byte(searchStr1), 16, 5))
	fmt.Println(searchStr2, "在", mainStr, "匹配的起始位置：", algorithm.BM([]byte(mainStr), []byte(searchStr2), 16, 3))
	fmt.Println(searchStr3, "在", mainStr, "匹配的起始位置：", algorithm.KMP([]byte(mainStr), []byte(searchStr3), 16, 3))
	fmt.Println(searchStr4, "在", mainStr, "匹配的起始位置：", algorithm.BM([]byte(mainStr), []byte(searchStr4), 16, 6))
}
