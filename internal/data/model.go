package data

import (
	"time"
)

type Order struct {
	OrderId   int64   `json:"order_id"`
	Uid       int64   `json:"uid"`
	Weight    float64 `json:"weight"`
	CreatedAt time.Time
}
