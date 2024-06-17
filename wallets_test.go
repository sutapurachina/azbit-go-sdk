package azbit_go_sdk

import (
	"fmt"
	"testing"
)

func TestAzBitClient_Balances(t *testing.T) {
	balances, err := azbit.Balances()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(balances)

	for _, balance := range balances.Balances {
		if balance.CurrencyCode == "BTC" {
			fmt.Println(balance.Amount)
		}
	}
}
