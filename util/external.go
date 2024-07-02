package util

import (
	"bytes"
	"eazyWallet/util/config"
	"eazyWallet/util/constant"
	"encoding/json"
	"io"
	"net/http"
)

func MakePostRequest[T any](url string, jsonData []byte, response T) (*T, error) {
	request, err := http.NewRequest(http.MethodPost, config.PaystackTransactionUrl, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	request.Header.Set(constant.ContentType, "application/json")
	request.Header.Set(constant.Authorization, url)

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	readAll, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(readAll, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
