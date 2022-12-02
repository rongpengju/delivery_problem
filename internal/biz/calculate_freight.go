package biz

import (
	"fmt"
	"github.com/rongpengju/delivery_problem/internal/pkg/etc"
	"math"

	"github.com/google/wire"
	"github.com/shopspring/decimal"
)

var ProviderSet = wire.NewSet(
	NewCourier,
)

type Courier struct {
	BasicPrice    int8
	OverPrice     int8
	BasicWeight   int8
	InsuranceRate int8
}

func NewCourier(config *etc.Config) *Courier {
	return &Courier{
		BasicPrice:    config.Calculate.BasicPrice,
		OverPrice:     config.Calculate.OverPrice,
		BasicWeight:   config.Calculate.BasicWeight,
		InsuranceRate: config.Calculate.InsuranceRate,
	}
}

func (c *Courier) Calculate(weight float64) float64 {
	if weight <= 0 || weight > 100 {
		return 0
	}

	//if weight <= 1 {
	//	return float64(c.BasicPrice)
	//}

	overPrice := c.overweightPrice(weight)
	insPrice := c.insurancePrice(weight)
	totalPrice := float64(c.BasicPrice) + float64(overPrice) + insPrice
	fmt.Printf("快递费明细: \n")
	fmt.Printf("\t基础费用: %v\n", c.BasicPrice)
	fmt.Printf("\t超重费用: %v\n", overPrice)
	fmt.Printf("\t保险费用: %v\n", insPrice)
	return round(totalPrice, 0)
}

// 计算超重费用
func (c *Courier) overweightPrice(weight float64) int {
	return int(math.Ceil(weight-float64(c.BasicWeight))) * int(c.OverPrice)
}

// 计算保险费用
func (c *Courier) insurancePrice(weight float64) float64 {
	var res float64
	w := int(math.Ceil(weight))
	for i := 1; i <= w; i++ {
		if i == w {
			break
		}
		res += float64(c.BasicPrice) + float64((i-int(c.BasicWeight))*int(c.OverPrice))
		res *= float64(c.InsuranceRate) * 0.01
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
