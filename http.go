package azbit_go_sdk

const baseApiUrl = "https://data.azbit.com"

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
