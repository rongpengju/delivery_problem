package gen

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateWeight() map[int][]float64 {
	rand.Seed(time.Now().UnixNano())
	return randFloats(0, 100, defaultUserNumber)
}

func randFloats(min, max float64, size int) map[int][]float64 {
	hash := make(map[int][]float64, 10)
	for i := 0; i < size; i++ {
		number, err := strconv.ParseFloat(fmt.Sprintf("%.2f", min+rand.Float64()*(max-min)), 64)
		if err != nil {
			return nil
		}
		switch {
		case number > 0 && number <= 10:
			hash[10] = append(hash[10], number)
		case number > 10 && number <= 20:
			hash[20] = append(hash[20], number)
		case number > 20 && number <= 30:
			hash[30] = append(hash[30], number)
		case number > 30 && number <= 40:
			hash[40] = append(hash[40], number)
		case number > 40 && number <= 50:
			hash[50] = append(hash[50], number)
		case number > 50 && number <= 60:
			hash[60] = append(hash[60], number)
		case number > 60 && number <= 70:
			hash[70] = append(hash[70], number)
		case number > 70 && number <= 80:
			hash[80] = append(hash[80], number)
		case number > 80 && number <= 90:
			hash[90] = append(hash[90], number)
		case number > 90 && number <= 100:
			hash[100] = append(hash[100], number)
		}
	}

	return hash
}
