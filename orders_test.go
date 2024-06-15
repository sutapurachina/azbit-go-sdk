package azbit_go_sdk

import (
	"fmt"
	"testing"
)

var azbit = NewAzBitClient("PSgZiaxoQ2GSJ2GrZI4MrNZbqUeLDkkTuQO6mA", "u2OSYn4ypAyo3VgT5hzEIunsfS9KV2YKyCnIu8I4Vw7CJcMSrk7vFevkhU68frXStTP8lQ")

func TestAzBitClient_OrderBook(t *testing.T) {
	levels, err := azbit.OrderBook("SNX", "USDT")
	if err != nil {
		fmt.Println(err)
	}

	for _, level := range levels {
		if level.IsBid {
			fmt.Println(level.Price)
		}
	}
}

func TestAzBitClient_MyOrders(t *testing.T) {
	orders, err := azbit.MyOrders("BTС", "USDT", "all")
	if err != nil {
		fmt.Println(err)
	}
	for _, order := range orders {
		fmt.Println(order)
	}
}

func TestAzBitClient_PostOrder(t *testing.T) {
	id, err := azbit.PostOrder(Buy, "BTC", "USDT", 0.000186, 59000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
}

func TestAzBitClient_CancelOrder(t *testing.T) {
	err := azbit.CancelOrder("cbd8bd76-1ba3-4b00-94c9-5cfc40406cfa")
	if err != nil {
		fmt.Println(err)
	}
}
