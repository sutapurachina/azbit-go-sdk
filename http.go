package azbit_go_sdk

import (
	"fmt"
	"io"
	"net/http"
)

const baseApiUrl = "https://data.azbit.com"

type response struct {
	Header     http.Header
	Body       io.ReadCloser
	StatusCode int
	Status     string
}

type KeyPair struct {
	publicKey string
	secretKey string
}

func (kp *KeyPair) Public() string {
	return kp.publicKey
}

func (kp *KeyPair) Secret() string {
	return kp.secretKey
}

func NewKeyPair(publicKey string, secretKey string) *KeyPair {
	return &KeyPair{publicKey: publicKey, secretKey: secretKey}
}

func checkHTTPStatus(resp *http.Response, expected ...int) error {
	for _, e := range expected {
		if resp.StatusCode == e {
			return nil
		}
	}
	return fmt.Errorf("http response status != %+v, got %d", expected, resp.StatusCode)
}
