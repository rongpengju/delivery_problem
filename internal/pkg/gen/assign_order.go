package gen

import (
	"math/rand"
	"time"
)

// AssignOrder 打散数据分配到100个用户
func AssignOrder() []int {
	rand.Seed(time.Now().UnixNano())
	randRes := randInt(900, 99, defaultUserNumber)
	var tmp int
	for _, v := range randRes {
		tmp += v
	}
	// 对剩下的数据再次处理
	average := (defaultOrderNumber - tmp) / defaultUserNumber
	for i, _ := range randRes {
		randRes[i] += average
	}
	randRes[0] += defaultOrderNumber - tmp - defaultUserNumber*average
	return randRes
}

func randInt(min, max int, size int) []int {
	numberList := make([]int, size)
	for i := 0; i < size; i++ {
		x := rand.Intn(max) + min
		numberList[i] = x
	}
	return numberList
}
