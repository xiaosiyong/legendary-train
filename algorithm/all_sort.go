package algorithm

import "fmt"

/***
// 数组全排列 调用方式：
// int[]a = a={1, 2, 3, 4}; printPermutations(a, 4, 4);
// k表示要处理的子数组的数据个数
*/
func PrintPermutations(data []int, n, k int) {
	if k == 1 {
		for i := 0; i < n; i++ {
			fmt.Print(data[i], " ")
		}
		fmt.Println()
	}
	for i := 0; i < k; i++ {
		data[i], data[k-1] = data[k-1], data[i]
		PrintPermutations(data, n, k-1)
		data[i], data[k-1] = data[k-1], data[i]
	}
}



func AllSort(d []int,n,k int){
	if k == 1 {
		for i := 0;i<n;i++{
			fmt.Print(d[i],"")
		}
		fmt.Println()
	}
	for i:=0;i<k;i++{
		d[i],d[k-1] = d[k-1],d[i]
		AllSort(d,n,k-1)
		d[i],d[k-1] = d[k-1],d[i]
	}

}
