package main

import (
	"fmt"

	"github.com/rongpengju/delivery_problem/internal/biz"
)

func main() {
	weight := 10.7
	fmt.Println(biz.Calculate(weight))
}
