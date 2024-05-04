package azbit_go_sdk

import (
	"fmt"
	"testing"
)

var azbit = NewAzBitClient("PSgZiaxoQ2GSJ2GrZI4MrNZbqUeLDkkTuQO6mA", "u2OSYn4ypAyo3VgT5hzEIunsfS9KV2YKyCnIu8I4Vw7CJcMSrk7vFevkhU68frXStTP8lQ")

func TestAzBitClient_OrderBook(t *testing.T) {
	levels, err := azbit.OrderBook("SHIB_USDT")
	if err != nil {
		fmt.Println(err)
	}

	for _, level := range levels {
		fmt.Println(level)
	}
}

func TestAzBitClient_MyOrders(t *testing.T) {
	orders, err := azbit.MyOrders("BTC", "USDT", "all")
	if err != nil {
		fmt.Println(err)
	}
	for _, order := range orders {
		fmt.Println(order)
	}
}

func TestAzBitClient_PostOrder(t *testing.T) {
	id, err := azbit.PostOrder(Buy, "BTC_USDT", 0.000186, 59000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
}

func TestAzBitClient_CancelOrder(t *testing.T) {
	err := azbit.CancelOrder("3db64569-5368-40a4-b257-6925f4dd1150")
	if err != nil {
		fmt.Println(err)
	}
}
