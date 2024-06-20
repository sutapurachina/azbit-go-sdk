package azbit_go_sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DealsRequest struct {
	CurrencyPairCode string
	SinceDate        string
	EndDate          string
	PageNumber       int
	PageSize         int
}

type Deal struct {
	ID               string  `json:"id"`
	DealDateUtc      string  `json:"dealDateUtc"`
	CurrencyPairCode string  `json:"currencyPairCode"`
	Volume           float64 `json:"volume"`
	Price            float64 `json:"price"`
	IsBuy            bool    `json:"isBuy"`
	IsUserBuyer      *bool   `json:"isUserBuyer"`
	OrderID          string  `json:"orderId"`
}

func (azbit *AzBitClient) Deals(request DealsRequest) (deals []Deal, err error) {
	url := baseApiUrl + "/api/deals"
	//query := fmt.Sprintf("?currencyPairCode=%s&sinceDate=%s&endDate=%s&pageNumber=%d&pageSize=%d", request.CurrencyPairCode, request.SinceDate,
	//	request.EndDate, request.PageNumber, request.PageSize)
	query := fmt.Sprintf("?currencyPairCode=%s", request.CurrencyPairCode)
	if request.SinceDate != "" {
		query += fmt.Sprintf("&sinceDate=%s", request.SinceDate)
	}
	if request.EndDate != "" {
		query += fmt.Sprintf("&endDate=%s", request.EndDate)
	}
	if request.PageNumber > 0 {
		query += fmt.Sprintf("&pageNumber=%d", request.PageNumber)
	}
	if request.PageSize > 0 {
		query += fmt.Sprintf("&pageSize=%d", request.PageSize)
	}
	req, err := http.NewRequest("GET", url+query, nil)
	if err != nil {
		return
	}
	resp, err := azbit.http.Do(req)
	if err != nil {
		return
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = checkHTTPStatus(resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("%v - %s, %s", err, resp.Status, string(respBody))
	}
	err = json.Unmarshal(respBody, &deals)
	return
}
