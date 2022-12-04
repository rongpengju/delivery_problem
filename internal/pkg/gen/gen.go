package gen

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rongpengju/delivery_problem/internal/data"
)

func GenerateAllOrder() []*data.Order {
	userId := GenerateUserId()
	orderId := GenerateID()
	weightInfo := GenerateWeight()
	assignOrderInfo := AssignOrder()

	// 每个重量所要生成的订单数量
	var hash = make(map[float64]int, len(weightInfo))
	for k, v := range weightInfo {
		for _, i := range v {
			_, ok := hash[i]
			if ok {
				hash[i] += int((float64(len(weightInfo[k])) * 0.01 * 100000) / float64(len(weightInfo[k])))
				continue
			}
			hash[i] = int((float64(len(weightInfo[k])) * 0.01 * 100000) / float64(len(weightInfo[k])))
		}
	}

	var allData = make([]*data.Order, 0)

	tmp := 0
	for i := 0; i < defaultUserNumber; i++ {
		for j := 0; j < assignOrderInfo[i]; j++ {
			// 生成订单ID并移除
			id := orderId[tmp]
			tmp += 1

			// 随机生成某个重量的 Order
		againRandKey:
			key := randWeightKey()
			weight := randWeightValue(weightInfo, key)
			// 如果被移除的无法 rand 出来则会重新 rand key，然后去 rand 其他的weight
			if weight == -1 {
				goto againRandKey
			}
			orderData := generateDatabaseOrder(id, int(userId[i]), weight)
			allData = append(allData, orderData)

			// 移除对 Weight的计数，如果不再有的时候，则从map中移除，不再进行 rand
			if !removeWeightCount(hash, weight) {
				weightInfo = removeWeightValue(key, weight, weightInfo)
			}
		}
	}

	fmt.Printf("本次创建数据: %d条\n", len(allData))
	return allData
}

func generateDatabaseOrder(id uint64, userId int, weight float64) *data.Order {
	order := &data.Order{
		OrderId:   int64(id),
		Uid:       int64(userId),
		Weight:    weight,
		CreatedAt: time.Now().Local(),
	}
	return order
}

func removeWeightCount(weightInfo map[float64]int, key float64) bool {
	value, _ := weightInfo[key]
	if value == 0 {
		return false
	}
	weightInfo[key] -= 1
	return true
}

func removeWeightValue(key int, value float64, weightInfo map[int][]float64) map[int][]float64 {
	for i, v := range weightInfo[key] {
		if v == value {
			weightInfo[key] = append(weightInfo[key][:i], weightInfo[key][i+1:]...)
			return weightInfo
		}
	}
	return weightInfo
}

func randWeightKey() int {
	weightKey := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	rand.Seed(time.Now().UnixNano())
	return weightKey[rand.Intn(len(weightKey))]
}

func randWeightValue(weightHash map[int][]float64, key int) float64 {
	value := weightHash[key]
	valueLength := len(value)
	if valueLength == 0 {
		return -1
	}
	n := rand.Intn(valueLength)
	return value[n]
}
