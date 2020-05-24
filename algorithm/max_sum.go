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

func GetMax(a [][]int, x int) int {
	w := len(a)
	//var  total int
	var t, m int
	if w > 0 {
		h := len(a[0])
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				if a[i][j] == x { //找到第一个点
					t++
				}
			}

		}
		//t为黑色总数
	}
	return m
}

func Step1(a [][]int) {
	w := len(a)
	t := 1
	if w > 0 {
		h := len(a[0])
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				if a[i][j] == 1 {
					a[i][j] = t
					t++
				}
			}

		}
		//t为黑色总数
	}
}

func Step2(a [][]int) {
	w := len(a)
	if w > 0 {
		h := len(a[0])
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				var t int
				if a[i][j] != 0 {
					t = a[i][j]
					if i-1 >= 0 && a[i-1][j] != 0 {
						if t >= a[i-1][j] {
							t = a[i-1][j]
						}
					}
					if j-1 >= 0 && a[i][j-1] != 0 {
						if t >= a[i][j-1] {
							t = a[i][j-1]
						}
					}
					if i+1 < w && a[i+1][j] != 0 {
						if t >= a[i+1][j] {
							t = a[i+1][j]
						}
					}
					if j+1 < h && a[i][j+1] != 0 {
						if t >= a[i][j+1] {
							t = a[i][j+1]
						}
					}
					a[i][j] = t
					if i-1 >= 0 && a[i-1][j] != 0 {

						a[i-1][j] = t

					}
					if j-1 >= 0 && a[i][j-1] != 0 {

						a[i][j-1] = t

					}
					if i+1 < w && a[i+1][j] != 0 {

						a[i+1][j] = t

					}
					if j+1 < h && a[i][j+1] != 0 {

						a[i][j+1] = t

					}
				}
			}

		}
		//t为黑色总数
	}
}

func Step3(a [][]int) int {
	w := len(a)

	m := make(map[int]int)
	if w > 0 {
		h := len(a[0])
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				if _, b := m[a[i][j]]; !b {
					m[a[i][j]] = 1
				} else {
					m[a[i][j]]++
				}
			}

		}

	}

	t := 1
	for k, v := range m {
		if k != 0 && v > t {
			t = v
		}
	}
	return t
}
