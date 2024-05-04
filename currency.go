package azbit_go_sdk

import (
	"encoding/json"
	"io"
	"net/http"
)

type CurrencyInfo struct {
	Code           string   `json:"code"`
	DigitsPrice    int      `json:"digitsPrice"`
	DigitsAmount   int      `json:"digitsAmount"`
	MinQuoteAmount *float64 `json:"minQuoteAmount"`
	MinBaseAmount  *float64 `json:"minBaseAmount"`
}

func (azbit *AzBitClient) Currencies() (currencies []CurrencyInfo, err error) {
	url := baseApiUrl + "/api/currencies/pairs"
	//signature := ab.Signature(url, "")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	resp, err := azbit.http.Do(req)
	if err != nil {
		return
	}
	respBody, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(respBody, &currencies)
	return
}
