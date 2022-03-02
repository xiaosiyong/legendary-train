package algorithm

import "math"

/*
回溯算法
*/

var (
	Items      = []int{2, 2, 4, 6, 3}
	ItemsValue = []int{3, 4, 8, 9, 6}
	Vol        = 9
	Count      = 5
	MaxV       = math.MinInt8
	_m         = [5][10]bool{}
)

//背包最大容量
func MaxVolumeInBag(i, cw int) {
	if i == Count || cw == Vol {
		if cw > MaxV {
			MaxV = cw
		}
		return
	}
	if ok := _m[i][cw]; ok {
		return
	}
	_m[i][cw] = true
	MaxVolumeInBag(i+1, cw) //不装第i个
	if cw+Items[i] <= Vol {
		MaxVolumeInBag(i+1, cw+Items[i])
	}
}

//背包最大值
func MaxValueInBag(i, cw, cv int) {
	if i == Count || cw == Vol {
		if cw > MaxV {
			MaxV = cw
		}
		return
	}
	if ok := _m[i][cw]; ok {
		return
	}
	_m[i][cw] = true
	MaxValueInBag(i+1, cw, cv) //不装第i个
	if cw+Items[i] <= Vol {
		MaxValueInBag(i+1, cw+Items[i], cv+ItemsValue[i])
	}
}
