package model

type Request struct {
	Id   string
	Data map[string]string
}

type TokenizeResponse struct {
	Id   string
	Data map[string]string
}

type TokenizedValue struct {
	Nonce           []byte
	CipherTextValue []byte
	Token           string
}

type DetokenizedResponse struct {
	Id   string
	Data map[string]DetokenizedValue
}

type DetokenizedValue struct {
	Found bool
	Value string
}
