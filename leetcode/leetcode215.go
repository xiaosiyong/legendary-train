package leetcode

import "math/rand"

//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//解法1、直接用快排  时间复杂度O(NLogN)
func FindKthLargest(nums []int, k int) int {
	//if len(nums) == 1 {
	//	return nums[0]
	//}
	//return partition3(nums,0,len(nums)-1,k)
	return findKthLargest(nums,k)

}
func partition3(nums []int,s,e,k int) int {
	for s < e {
		a := nums[s]
		p := e
		for i := e;i>0;i--{
			if nums[i] < a {
				nums[i],nums[p] = nums[p],nums[i]
				p--
			}
		}
		nums[p],nums[s] = nums[s],nums[p]
		if p == k-1 {
			break
		}else if  p > k-1 {
			e = p-1

		}else {
			s = p+1
		}
	}
	return nums[k-1]
}

func findKthLargest(nums []int, k int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	return quickSelect(nums,0,l-1,k)

}


func partition4(nums []int,s,e,r int) int {
		a := nums[r]
		nums[r],nums[s] = nums[s],nums[r]
		p := e
		for i := e;i>0;i--{
			if nums[i] < a {
				nums[i],nums[p] = nums[p],nums[i]
				p--
			}
		}
		nums[p],nums[s] = nums[s],nums[p]
		return p

}

func quickSelect(nums []int,s,e,k int)int{
	if s == e {
		return nums[s]
	}
	r := s + rand.Intn(e-s)
	p := partition4(nums,s,e,r)
	if p == k-1 {
		return nums[p]
	} else if p > k-1 {
		return quickSelect(nums,0,p-1,k)
	}else {
		return quickSelect(nums,p+1,e,k)
	}

}








