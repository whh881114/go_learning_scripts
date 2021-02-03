package main

import "fmt"

type stockPosition struct {
	tickerSymbol string
	sharePrice   float32
	count        float32
}

// 获取stock的值（价值）
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	brand string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("资产的价值是：%f\n", asset.getValue())
}

func main() {
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)

	o = car{"BMW", "M3", 66500}
	showValue(o)

}
