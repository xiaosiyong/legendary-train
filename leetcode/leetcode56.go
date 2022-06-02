package leetcode

import s "sort"

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
[[1,5],[2,6],[4,8]]
[3][2]int
*/

func mergeSlice(intervals [][]int) [][]int {
	l := len(intervals)
	if l <= 1 {
		return intervals
	}
	s.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//注意原切片长度会变，所以之类不能用l作为长度
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i][1] = max(intervals[i][1], intervals[i+1][1])
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			i--
		}
	}
	return intervals
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
