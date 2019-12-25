package leetcode

import "fmt"

/**
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
 */
func ThreeSum(nums []int,sum int)[][]int{
	if len(nums) < 3 {
		return nil
	}
	var result [][]int
	var sets [][3]int
	for i := 0;i < len(nums);i++ {
		for j := i+1;j<len(nums);j++{
			for t := j+1;t<len(nums);t++{
				if nums[i] + nums[j] + nums[t] == sum {
					temp := sort(nums[i],nums[j],nums[t])
					if !checkExist(temp,sets) {
						result = append(result,temp)
						var m [3]int
						m[0] = temp[0]
						m[1] =temp[1]
						m[2] =temp[2]
						sets = append(sets,m)
					}
				}
			}
		}
	}
	return result
}

func checkExist(a []int,b [][3]int)bool {
	if  len(b)<1{
		return false
	}
	for i := 0;i<len(b);i++{
		if a[0]==b[i][0] && a[1]==b[i][1]&&a[2]==b[i][2] {
			return true
		}

	}
	return false
}

func sort(a,b,c int)[]int{
	var temp []int
	if a >= b && b >= c {
		temp = append(temp,c)
		temp = append(temp,b)
		temp = append(temp,a)
	}else 	if a >= c && c >= b {
		temp = append(temp,b)
		temp = append(temp,c)
		temp = append(temp,a)
	} else 	if b >= c && c >= a {
		temp = append(temp,a)
		temp = append(temp,c)
		temp = append(temp,b)
	} else 	if b >= a && a >= c {
		temp = append(temp,c)
		temp = append(temp,a)
		temp = append(temp,b)
	}else 	if c >= b && b >= a {
		temp = append(temp,a)
		temp = append(temp,b)
		temp = append(temp,c)
	}else {
		temp = append(temp,b)
		temp = append(temp,a)
		temp = append(temp,c)
	}
	return  temp
}

func ThreeSum2(nums []int)[][]int{
	var result [][]int
	if len(nums)<3 {
		return result
	}
	quickSort(nums,0,len(nums)-1)
	fmt.Println(nums)
	for i := 0;i < len(nums);i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		h := len(nums)-1
		for l < h {
			sum := nums[i] + nums[l] + nums[h]
			if sum == 0 {
				var temp []int
				temp = append(temp,nums[i])
				temp = append(temp,nums[l])
				temp = append(temp,nums[h])
				result = append(result,temp)
				for l < h && nums[l] == nums[l+1]{
					l++
				}
				for l < h && nums[h] == nums[h-1]{
					h--
				}
				l++
				h--
			}else if sum < 0 {
				l++
			}else{
				h--
			}
		}


	}
	return result
}


//快速排序
func quickSort(nums []int,l,h int){
	if l >= h {
		return
	}
	m := partition2(nums,l,h)
	quickSort(nums,l,m-1)
	quickSort(nums,m+1,h)
}

func partition(nums []int,l,h int)int{
	cursor := nums[h]
	j := l
	for i:=l;i<h;i++{
		if nums[i] < cursor {
			nums[i],nums[j] = nums[j],nums[i]
			j++
		}
	}
	nums[j],nums[h] =nums[h],nums[j]
	return j
}

func partition2(nums []int,l,h int)int{
	cursor := nums[h]
	j := h
	for i:=l;i<h;i++{
		if nums[i] > cursor {
			nums[i],nums[j] = nums[j],nums[i]
			j--
		}
	}
	nums[j],nums[h] =nums[h],nums[j]
	return j
}
