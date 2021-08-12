package algorithm

//最大子序和

func GetMaxSum(array []int) int {
	if array != nil {
		if len(array) < 2 {
			return array[0]
		}
		var max, temp int
		for i := 0; i < len(array); i++ {
			temp += array[i]
			if temp > max {
				max = temp
			} else if temp < 0 {
				temp = 0
			}
		}
		return max
	}
	return 0
}

func GetMaxSum2(array []int) int {
	if array != nil {
		if len(array) < 2 {
			return array[0]
		}
		var sum int
		max := array[0]
		for i := 0; i < len(array); i++ {
			sum += array[i]
			if sum < 0 {
				if max < sum {
					max = sum
				}
				sum = 0
			}
		}
		if sum > max && sum > 0 {
			return sum
		}
		return max
	}
	return 0
}


func IsExistNSum(a []int,m int)bool{
	if a == nil {
		return false
	}
	return false

}

func sum(a []int,n,m int)bool{
	if n <= 0 || m <=0 {
		return false
	}
	return false
}