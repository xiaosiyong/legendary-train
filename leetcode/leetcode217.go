package leetcode

/***
给定一个整数数组，判断是否存在重复元素。
如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
*/
func containsDuplicate(nums []int) bool {
	l := len(nums)
	if l > 0 {
		m := make(map[int]int, l)
		for i := 0; i < l; i++ {
			if _, b := m[nums[i]]; !b {
				m[nums[i]] = 1
			} else {
				return true
			}
		}
	}
	return false
}

/***
给定一个数组nums和一个整数k，是否存在两个不相等的整数 i 和 j，使得nums[i] == nums[j]，并且i和j之间的距离最多为k。
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	l := len(nums)
	if l > 0 {
		m := make(map[int]int, l)
		for i := 0; i < l; i++ {
			if _, b := m[nums[i]]; !b {
				m[nums[i]] = i
			} else {
				if i-m[nums[i]] <= k {
					return true
				} else {
					m[nums[i]] = i
				}

			}
		}
	}
	return false
}

/***
给定整数数组，是否存在两个不相等的 i 和 j 使得nums[i] 和nums[j] 的差值至多为 t ，i 和 j的距离至多为k。
*/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	l := len(nums)
	if l > 0 {

		m := make(map[int]int, l)
		for i := 0; i < l; i++ {
			if _, b := m[nums[i]]; !b {
				m[nums[i]] = i
			} else {
				if i-m[nums[i]] <= k {
					return true
				} else {
					m[nums[i]] = i
				}

			}
		}
	}
	return false
}
