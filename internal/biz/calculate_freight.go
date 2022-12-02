package biz

import (
	"math"

	"github.com/shopspring/decimal"
)

const (
	basicWeight = 1

	defaultBasicPrice int = 18

	defaultOverPrice int = 5

	insuranceRate = 0.01
)

func Calculate(weight float64) float64 {
	if weight <= 0 || weight > 100 {
		return 0
	}

	if weight <= 1 {
		return float64(defaultBasicPrice)
	}

	overPrice := overweightPrice(weight)
	insPrice := insurancePrice(weight)
	totalPrice := float64(defaultBasicPrice) + float64(overPrice) + insPrice
	return round(totalPrice, 0)
}

// 计算超重的费用
func overweightPrice(weight float64) int {
	return int(math.Ceil(weight-basicWeight)) * defaultOverPrice
}

// 计算保险费用
func insurancePrice(weight float64) float64 {
	var res float64
	w := int(math.Ceil(weight))
	for i := 1; i <= w; i++ {
		if i == w {
			break
		}
		res += float64(defaultBasicPrice) + float64((i-basicWeight)*defaultOverPrice)
		res *= insuranceRate
	}
	return processFloat(res)
}

// 处理保险费用的小数点，从小数点第三位做四舍五入处理
func processFloat(price float64) float64 {
	value, _ := decimal.NewFromFloat(price).Round(2).Float64()
	return value
}

// 对整体价格进行四舍五入的处理
func round(price float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(price*p+0.5) / p
}
