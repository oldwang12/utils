package easyhttp

import (
	"bytes"
	"io"
	"net/http"
)

func Request(url, method string, requestBody []byte, headers map[string]string) ([]byte, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(resp.Body)
}
