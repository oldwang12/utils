package easyhttp

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
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

// 分享内容中中提取链接
func ExtractURL(text string) (string, error) {
	// Define the regular expression pattern for matching URLs
	urlPattern := `https?://[^\s]+`

	// Compile the regular expression
	re := regexp.MustCompile(urlPattern)

	// Find the first match in the text
	match := re.FindString(text)

	if match == "" {
		return "", fmt.Errorf("no url found in the provided text")
	}
	return match, nil
}
