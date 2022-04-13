package models

import "container/list"

type CarInventory struct {
	List       list.List //库存队列
	SellCount  int
	TotalCount int
	Rate       float64
	X          float64
}

// AgeModel age
type AgeModel struct {
	Age int `json:"age"` //年龄
}
