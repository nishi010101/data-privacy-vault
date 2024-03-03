package model

type Request struct {
	Id   string            `json:"id,omitempty"`
	Data map[string]string `json:"data,omitempty"`
}

type TokenizeResponse struct {
	Id   string            `json:"id,omitempty"`
	Data map[string]string `json:"data,omitempty"`
}

type TokenizedValue struct {
	Nonce           []byte
	CipherTextValue []byte
	Token           string
}

type DetokenizedResponse struct {
	Id   string                      `json:"id,omitempty"`
	Data map[string]DetokenizedValue `json:"data,omitempty"`
}

type DetokenizedValue struct {
	Found bool   `json:"found,omitempty"`
	Value string `json:"value,omitempty"`
}
