package leetcode

/*整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。*/
func Search(nums []int, target int) int {
	if nums == nil {
		return -1
	}
	l := len(nums)
	if l == 0 {
		return -1
	}
	if l == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	r := findRotateIndex(nums, 0, l-1)
	if nums[r] == target {
		return r
	}
	if r == 0 {
		return bSearch(nums, 0, l-1, target)
	}
	if nums[0] > target {
		return bSearch(nums, r+1, l-1, target)
	} else {
		return bSearch(nums, 0, r-1, target)
	}
	return -1
}

func findRotateIndex(nums []int, l, r int) int {
	if nums[l] < nums[r] {
		return 0
	}
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] > nums[m+1] {
			return m + 1
		} else {
			if nums[m] < nums[l] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return 0
}

func bSearch(nums []int, l, r, t int) int {
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] == t {
			return m
		} else {
			if nums[m] > t {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}
