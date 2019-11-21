package AlgorithmLearn

//字符串匹配算法

//用来存放散列表的Hash
var size = 256

func BM(mainBytes []byte, searchBytes []byte, mainL int, searchL int) int {
	index := -1
	if searchL >= mainL {
		index = -1
	}
	hash := [256]int{}            //记录模式串中每个字符最后出现的位置
	generateBc(searchBytes, hash) //生成坏字符哈希表
	var i int                     //i表示主串与模式串对齐的第一个字符

	suffix := make([]int, mainL)
	prefix := make([]bool, mainL)
	generateGS(searchBytes, suffix, prefix)

	for i <= mainL-searchL {
		var j int
		for j = searchL - 1; j >= 0; j-- { //模式串从后往前匹配
			if mainBytes[i+j] != searchBytes[j] { //坏字符对应模式串中的下标是j
				break
			}
		}
		if j < 0 {
			return i // 匹配成功，返回主串与模式串第一个匹配的字符的位置
		}
		x := j - hash[int(mainBytes[i+j])]
		y := 0
		if j < mainL-1 { //如果有好后缀
			y = moveByGs(j, mainL, suffix, prefix)
		}
		if x > y {
			i = i + x
		} else {
			i = i + y
		}
	}

	return index
}

//预处理串
func generateBc(searchBytes []byte, hash [256]int) {
	for i := 0; i < size; i++ {
		hash[i] = -1
	}
	for i := 0; i < len(searchBytes); i++ {
		ascii := int(searchBytes[i])
		hash[ascii] = i
	}
}

func generateGS(searchBytes []byte, suffix []int, prefix []bool) {
	m := len(searchBytes)
	for i := 0; i < m; i++ {
		suffix[i] = -1
		prefix[i] = false
	}
	for i := 0; i < m-1; i++ {
		j := i
		var k int
		for j >= 0 && searchBytes[j] == searchBytes[m-1-k] {
			j--
			k++
			suffix[k] = j + 1
		}
		if j == -1 {
			prefix[k] = true
		}
	}
}

func moveByGs(j, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j //好后缀的长度
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}
	for r := j + 2; r <= m-1; r++ {
		if prefix[m-r] {
			return r
		}
	}
	return m
}

//BF字符串匹配算法
func BF(mainBytes []byte, searchBytes []byte) (index int) {
	return index
}

//RK字符串匹配算法
func RK(mainBytes []byte, searchBytes []byte) (index int) {
	return index
}


func KMP(m []byte,s []byte,ml,sl int)int{
	nexts := kmpGetNext(s,sl)
	j := 0
	for i := 0; i < ml; i++{
		for j > 0 && m[i]!=s[j]{//一直找到a[i]和b[j]
			j = nexts[j - 1] + 1
		}
		if m[i] == s[j]{
			j++
		}
		if j == sl {
			return i - sl +1
		}
	}
	return -1
}

func kmpGetNext(s []byte,sl int)[]int{
	next := make([]int,sl)
	next[0] = -1
	k := -1
	for i := 1; i < sl; i++ {
		for k != -1 && s[k+1]!= s[i] {
			k = next[k]
		}
		if s[k + 1] == s[i]{
			k++
		}
		next[i] = k
	}
	return next
}