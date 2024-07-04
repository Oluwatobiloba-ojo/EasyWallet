package util

import (
	"encoding/base64"
	"github.com/google/uuid"
)

type Encoding struct {
}

func NewEncoding() *Encoding {
	return &Encoding{}
}

func (encode Encoding) EncodeTo(byte []byte) string {
	encoding := base64.StdEncoding
	return encoding.EncodeToString(byte)
}

func GenerateRefrenceCode() string {
	return uuid.New().String()
}
