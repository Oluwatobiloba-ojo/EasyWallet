package util

import (
	"bytes"
	"eazyWallet/util/constant"
	"eazyWallet/util/errorMessage"
	"encoding/json"
	"io"
	"net/http"
)

func MakePostRequest[T any](key string, jsonData []byte, response T, url string) (*T, error) {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	request.Header.Set(constant.ContentType, constant.APPLICATION_JSON)
	request.Header.Set(constant.Authorization, key)

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	} else if res.StatusCode == http.StatusOK {
		return extractResponse[T](res, response)
	} else {
		return nil, errorMessage.PaymentTransactionFailed()
	}
}

func extractResponse[T any](res *http.Response, response T) (*T, error) {
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
