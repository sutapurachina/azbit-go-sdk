package azbit_go_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Side string

var (
	Buy  Side = "buy"
	Sell Side = "sell"
)

type PostOrderRequest struct {
	OrderSide        Side    `json:"side"`
	CurrencyPairCode string  `json:"currencyPairCode"`
	Amount           float64 `json:"amount"`
	Price            float64 `json:"price"`
}

type PostOrderData struct {
	ID string `json:"id"`
}

type PostOrderResponse struct {
	Data       PostOrderData `json:"data"`
	Succeeded  bool          `json:"succeeded"`
	Errors     []string      `json:"errors"`
	LogicError string        `json:"logicError"`
}

type CancelOrderResponse struct {
	Data       PostOrderData `json:"data"`
	Succeeded  bool          `json:"succeeded"`
	Errors     []string      `json:"errors"`
	LogicError string        `json:"logicError"`
}

type BookLevel struct {
	IsBid        bool    `json:"isBid"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	CurrencyTo   float64 `json:"currencyTo"`
	QuoteAmount  float64 `json:"quoteAmount"`
	CurrencyFrom float64 `json:"currencyFrom"`
}

type Order struct {
	ID               string  `json:"id"`
	IsBid            bool    `json:"isBid"`
	Price            float64 `json:"price"`
	InitialAmount    float64 `json:"initialAmount"`
	Amount           float64 `json:"amount"`
	CurrencyTo       float64 `json:"currencyTo"`
	QuoteAmount      float64 `json:"quoteAmount"`
	CurrencyFrom     float64 `json:"currencyFrom"`
	Date             string  `json:"date"`
	UserID           string  `json:"userId"`
	IsCanceled       bool    `json:"isCanceled"`
	Status           string  `json:"status"`
	CurrencyPairCode string  `json:"currencyPairCode"`
}

func (azbit *AzBitClient) OrderBook(base, quote string) (levels []BookLevel, err error) {
	url := baseApiUrl + "/api/orderbook/"
	currencyPairCode := symbol(base, quote)
	query := fmt.Sprintf("?currencyPairCode=%s", currencyPairCode)
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
		return
	}
	err = json.Unmarshal(respBody, &levels)
	return
}

func (azbit *AzBitClient) MyOrders(base, quote, status string) (orders []Order, err error) {
	url := baseApiUrl + "/api/user/orders"
	query := fmt.Sprintf("?currencyPairCode=%s&status=%s", symbol(base, quote), status)
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
	err = json.Unmarshal(respBody, &orders)
	return
}

func symbol(base, quote string) string {
	return base + "_" + quote
}

func (azbit *AzBitClient) PostOrder(side Side, base, quote string, amount, price float64) (ID string, err error) {
	return azbit.postOrder(
		&PostOrderRequest{
			OrderSide:        side,
			CurrencyPairCode: symbol(base, quote),
			Amount:           amount,
			Price:            price,
		})
}

func (azbit *AzBitClient) postOrder(postOrderRequest *PostOrderRequest) (ID string, err error) {
	url := baseApiUrl + "/api/v200/orders"
	query := ""
	reqBodyJson, err := json.Marshal(postOrderRequest)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", url+query, bytes.NewReader(reqBodyJson))
	if err != nil {
		return
	}
	azbit.sign(req, url+query, string(reqBodyJson))
	resp, err := azbit.http.Do(req)
	if err != nil {
		return
	}
	respBody, err := io.ReadAll(resp.Body)
	response := PostOrderResponse{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return
	}
	if response.Succeeded {

		ID = response.Data.ID
		if ID == "" {
			err = errors.New("didn't get the order id. seems like order is not placed")
		}
		return
	}
	errorStr := ""
	for _, errStr := range response.Errors {
		errorStr += errStr + " "
	}
	errorStr += response.LogicError
	err = errors.New(errorStr)
	return
}

func (azbit *AzBitClient) CancelOrder(orderId string) (err error) {
	url := baseApiUrl + "/api/v200/orders/"
	req, err := http.NewRequest("DELETE", url+orderId, nil)
	if err != nil {
		return
	}
	azbit.sign(req, url+orderId, "")
	resp, err := azbit.http.Do(req)
	if err != nil {
		return
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	response := CancelOrderResponse{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return
	}

	if response.Succeeded {
		return nil
	}
	errorStr := ""
	for _, errStr := range response.Errors {
		errorStr += errStr + " "
	}
	errorStr += response.LogicError
	err = errors.New(errorStr)
	return
}
