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
