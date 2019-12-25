package leetcode

func Search(nums []int, target int) int {
	if nums == nil{
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
	r := findRotateIndex(nums,0,l-1)
	if nums[r] == target {
		return r
	}
	if r == 0 {
		return bSearch(nums,0,l-1,target)
	}
	if nums[0] > target {
		return bSearch(nums,r+1,l-1,target)
	}else{
		return bSearch(nums,0,r-1,target)
	}
	return -1
}

func findRotateIndex(nums []int,l,r int)int{
	if nums[l] < nums[r] {
		return 0
	}
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] > nums[m+1]{
			return m+1
		}else{
			if nums[m] < nums[l] {
				r = m -1
			}else{
				l = m + 1
			}
		}
	}
	return 0
}

func bSearch(nums []int, l,r,t int)int{
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] == t {
			return m
		}else {
			if nums[m] > t {
				r = m - 1
			}else{
				l = m + 1
			}
		}
	}
	return -1
}