package azbit_go_sdk

import (
	"fmt"
	"testing"
)

func TestAzBitClient_Deals(t *testing.T) {
	de, err := azbit.Deals(DealsRequest{CurrencyPairCode: "BTC_USDT",
		SinceDate: "2024-06-18", EndDate: "2024-06-19", PageSize: 200, PageNumber: 1})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(de))
	for _, d := range de {
		fmt.Println(d)
	}
}
