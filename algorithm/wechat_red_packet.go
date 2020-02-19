package algorithm

import (
	"math"
	"math/rand"
	"time"
)

type RedPacket struct {
	TotalQuantity float64
	TotalMoney float64
	RemainQuantity float64
	RemainMoney float64
}

func RandomMoney(p *RedPacket)float64{
	var r,min float64
	if p.RemainQuantity == 1{
		return math.Round(p.RemainMoney*100)/100
	}
	min = 0.01
	max := p.RemainMoney / p.RemainQuantity *2
	rand.Seed(time.Now().UnixNano())
	r = rand.Float64() * max
	if r < min {
		r = min
	}
	r = math.Floor(r * 100)/100
	p.RemainQuantity--
	p.RemainMoney -= r
	return r
}
