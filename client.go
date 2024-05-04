package azbit_go_sdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

type AzBitClient struct {
	keyPair *KeyPair
	http    *http.Client
}

func NewAzBitClient(publicKey string, secretKey string) *AzBitClient {
	client := http.DefaultClient
	return newAzBitClient(NewKeyPair(publicKey, secretKey), client)
}

func newAzBitClient(keyPair *KeyPair, client *http.Client) *AzBitClient {
	return &AzBitClient{
		keyPair: keyPair,
		http:    client,
	}
}

func (azbit *AzBitClient) signature(requestUrl, requestBodyString string) string {
	signatureText := azbit.keyPair.Public() + requestUrl + requestBodyString
	// Convert signatureText to bytes
	signatureBytes := []byte(signatureText)

	// Compute HMACSHA256 hash
	h := hmac.New(sha256.New, []byte(azbit.keyPair.Secret()))
	h.Write(signatureBytes)
	hash := h.Sum(nil)

	// Convert hash to HEX-string
	hexHash := hex.EncodeToString(hash)

	return hexHash
}

func (azbit *AzBitClient) sign(req *http.Request, requestUrl, requestBodyString string) {
	if requestBodyString != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("API-PublicKey", azbit.keyPair.Public())
	req.Header.Set("API-Signature", azbit.signature(requestUrl, requestBodyString))
}
