package leetcode

import (
	"legendaryTrain/datastruct"
	"math"
	"strconv"
)

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
func LargestRectangleArea1(heights []int) int {
	h := len(heights)
	if h > 0 {
		maxArea := 0
		for i := 0; i < h; i++ {
			for j := i; j < h; j++ {
				m := heights[j]
				for k := i; k <= j; k++ {
					if heights[k] < m {
						m = heights[k]
					}
				}
				if (j-i+1)*m > maxArea {
					maxArea = (j - i + 1) * m
				}
			}
		}
		return maxArea
	}
	return 0
}

func LargestRectangleArea2(h []int) int {
	if len(h) > 0 {
		maxArea := 0
		for i := 0; i < len(h); i++ {
			m := h[i]
			for j := i; j < len(h); j++ {
				if h[j] < m {
					m = h[j]
				}
				if (j-i+1)*m > maxArea {
					maxArea = (j - i + 1) * m
				}
			}
		}
		return maxArea
	}
	return 0
}

func LargestRectangleArea3(h []int) int {
	return calculateArea(h, 0, len(h)-1)
}

func calculateArea(h []int, s, e int) int {
	if s > e {
		return 0
	}
	minIndex := s
	for i := s; i <= e; i++ {
		if h[i] < h[minIndex] {
			minIndex = i
		}
	}
	return int(math.Max(float64((e-s+1)*h[minIndex]), math.Max(float64(calculateArea(h, 0, minIndex-1)), float64(calculateArea(h, minIndex+1, e)))))

}

func LargestRectangleArea4(h []int) int {
	if len(h) <= 0 {
		return 0
	}
	s := datastruct.InitArrayStack(len(h) + 1)
	s.Push("-1")
	h = append(h, -1)
	a := 0
	for i := 0; i < len(h); i++ {
		for s.Peek() != "-1" && h[i] < h[stringToInt(s.Peek())] {
			t := h[stringToInt(s.Peek())]
			s.Pop()
			if t*(i-stringToInt(s.Peek())-1) > a {
				a = t * (i - stringToInt(s.Peek()) - 1)
			}
		}
		s.Push(strconv.Itoa(i))
	}
	return a
}

func stringToInt(s string) int {
	if i, e := strconv.Atoi(s); e == nil {
		return i
	}
	return -1
}
