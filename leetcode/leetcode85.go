package leetcode

/***
给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
*/
func MaximalRectangle1(m [][]int) int {
	//先计算每个点所在位置的最大宽度
	for i := 0; i < len(m); i++ {
		for j := 1; j < len(m[i]); j++ {
			if m[i][j] == 1 && m[i][j-1] == 1 {
				m[i][j] = m[i][j-1] + 1
			}
		}
	}
	a := 0
	//纵向遍历取宽度最小的 作为宽  纵向的跨度作为高  计算面积
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			w := m[i][j]
			for k := i; k > 0; k-- {
				if m[i][j] > m[k][j] {
					w = m[k][j]
				}
				if w*(i-k+1) > a {
					a = w * (i - k + 1)
				}
			}
		}
	}
	return a
}

func MaximalRectangle2(m [][]int) int {
	var max int
	if len(m) > 0 {
		t := make([]int, len(m[0]))
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[i]); j++ {
				if m[i][j] == 1 {
					t[j]++
				} else {
					t[j] = 0
				}
			}
			c := LargestRectangleArea4(t)
			if c > max {
				max = c
			}
		}
	}
	return max
}

func MaximalRectangle3(s [][]int) int {
	var a int
	if len(s) > 0 {
		m := len(s)
		n := len(s[0])
		l, r, h := make([]int, n), make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			r[i] = n
			l[i] = -1
		}
		for i := 0; i < m; i++ {

			//update height
			for j := 0; j < n; j++ {
				if s[i][j] == 1 {
					h[j] += 1
				} else {
					h[j] = 0
				}

			}
			//update left
			cl := -1
			for j := 0; j < n; j++ {
				if s[i][j] == 1 {
					if l[j] < cl {
						l[j] = cl
					}
				} else {
					l[j] = -1
					cl = j
				}
			}
			//update right
			cr := n
			for j := n - 1; j >= 0; j-- {
				if s[i][j] == 1 {
					if r[j] > cr {
						r[j] = cr
					}
				} else {
					r[j] = n
					cr = j
				}
			}
			//update area
			for j := 0; j < n; j++ {
				t := (r[j] - l[j] - 1) * h[j]
				if a < t {
					a = t
				}
			}

		}
	}
	return a
}
