package azbit_go_sdk

import (
	"fmt"
	"testing"
)

func TestAzBitClient_Currencies(t *testing.T) {
	azbit := NewAzBitClient("", "")
	currencies, err := azbit.Currencies()
	if err != nil {
		fmt.Println(err)
	}
	for _, currency := range currencies {
		fmt.Println(currency)
	}
}
