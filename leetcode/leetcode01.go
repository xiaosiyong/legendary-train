package leetcode

import "legendaryTrain/algorithm"

func twoSum(nums []int, target int) []int {
	var r []int
	if len(nums) > 1 {
		t := len(nums)
		m := make(map[int]int, t)
		for i := 0; i < t; i++ {
			a := target - nums[i]
			b, ok := m[a]
			if ok && b != i {
				r = append(r, i)
				r = append(r, b)
				break
			} else {
				m[nums[i]] = i
			}
		}
	}
	return r
}

func TwoSum2(nums []int, target int) []int {
	var r []int
	h := len(nums)
	if h > 1 {
		algorithm.QuickSort(nums, 0, h-1)
		i := 0
		j := h - 1
		for i < j {
			if nums[i]+nums[j] == target {
				r = append(r, i)
				r = append(r, j)
				return r
			} else if nums[i]+nums[j] > target {
				j--
			} else {
				i++
			}
		}
	}
	return r
}
