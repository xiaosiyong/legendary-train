package algorithm

/***
0 ，1矩阵中含有连续全1块的个数
一个矩阵中只有0和1两种值， 每个位置都可以和自己的上、 下、 左、 右
四个位置相连， 如果有一片1连在一起， 这个部分叫做一个岛， 求一个
矩阵中有多少个岛？
*/
func GetIslandQuantity(m [][]int) int {
	var q int
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 1 {
				q++
				dfIsland(m, i, j)
			}
		}
	}
	return q
}

func dfIsland(m [][]int, i, j int) {
	m[i][j] = 0
	if i-1 >= 0 && m[i-1][j] == 1 {
		dfIsland(m, i-1, j)
	}
	if j-1 >= 0 && m[i][j-1] == 1 {
		dfIsland(m, i, j-1)
	}
	if i+1 < len(m) && m[i+1][j] == 1 {
		dfIsland(m, i+1, j)
	}
	if j+1 < len(m[0]) && m[i][j+1] == 1 {
		dfIsland(m, i, j+1)
	}
}

/***
对矩阵做三次扫描, 扫描的次序都是从左到右,从上到下.第一遍将所有为1的元素,依次标一个值,这个值从1开始 例如:
0 0 0 0 0 // 0 0 0 0 0标注成这样
0 1 1 0 0 // 0 1 2 0 0
0 1 0 0 0 // 0 3 0 0 0
0 0 0 1 0 // 0 0 0 4 0
0 0 0 1 0 // 0 0 0 5 0
第二遍 如果 遇到不为0 的元素, 考虑这个元素以及它周围四个元素中,不为0的元素标，求他们的最小值，并将这些元素通标注为这个最小值　像上面，遇到了１　那么将 2和3 标注成 1 。就变成下面这样
0 0 0 0 0
0 1 1 0 0
0 1 0 0 0
0 0 0 4 0
0 0 0 4 0
最后一遍 数一下各种数字的个数，最多的，就是最大全1子块的面积。(注，这样做，结果是连续的全1区域，很可能是不规则的)
*/
//连续全1的个数

func GetContinuousQuantity(m [][]int) int {
	var q int
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 1 {
				var t int
				df(m, i, j, &t)
				if t > q {
					q = t
				}
			}
		}
	}
	return q
}

func df(m [][]int, i, j int, q *int) {
	*q++
	m[i][j] = 0
	if i-1 >= 0 && m[i-1][j] == 1 {
		df(m, i-1, j, q)
	}
	if j-1 >= 0 && m[i][j-1] == 1 {
		df(m, i, j-1, q)
	}
	if i+1 < len(m) && m[i+1][j] == 1 {
		df(m, i+1, j, q)
	}
	if j+1 < len(m[0]) && m[i][j+1] == 1 {
		df(m, i, j+1, q)
	}

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
