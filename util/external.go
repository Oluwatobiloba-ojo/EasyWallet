package util

import (
	"bytes"
	"eazyWallet/util/constant"
	"eazyWallet/util/message"
	"encoding/json"
	"io"
	"log"
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
		return ExtractResponse[T](res, response)
	} else {
		log.Println(ExtractResponse[T](res, response))
		return nil, message.PaymentTransactionFailed()
	}
}

func ExtractResponse[T any](res *http.Response, response T) (*T, error) {
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
