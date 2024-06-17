package azbit_go_sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Balance struct {
	Amount              float64 `json:"amount"`
	AmountBtc           float64 `json:"amountBtc"`
	AmountUsdt          float64 `json:"amountUsdt"`
	AmountInternalToken float64 `json:"amountInternalToken"`
	CurrencyCode        string  `json:"currencyCode"`
	CurrencyName        string  `json:"currencyName"`
	Digits              int     `json:"digits"`
	CurrencyIsFiat      bool    `json:"currencyIsFiat"`
}

type BalancesResponse struct {
	Balances []Balance `json:"balances"`
}

func (azbit *AzBitClient) Balances() (balances *BalancesResponse, err error) {
	url := baseApiUrl + "/api/wallets/balances"
	query := ""
	req, err := http.NewRequest("GET", url+query, nil)
	if err != nil {
		return
	}
	azbit.sign(req, url+query, "")
	resp, err := azbit.http.Do(req)
	if err != nil {
		return
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = checkHTTPStatus(resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("%v - %s, %s", err, resp.Status, string(respBody))
	}

	err = json.Unmarshal(respBody, &balances)
	return
}
