package easyhttp

import (
	"bytes"
	"net/http"
)

func Request(url, method string, requestBody []byte,headers map[string]string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}
	client := &http.Client{}
	return client.Do(request)
}