package common

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func GetDefaultHeaders() map[string]string {
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "*/*"
	headers["Connection"] = "keep-alive"
	headers["Accept-Language"] = "zh-CN,zh;q=0.9"
	return headers
}

func HttpRequest(url, method string, requestBody []byte, headers map[string]string) ([]byte, error) {
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

// 分享内容中提取链接
func ExtractURL(text string) (string, error) {
	// Define the regular expression pattern for matching URLs
	urlPattern := `https?://[^\s]+`

	// Compile the regular expression
	re := regexp.MustCompile(urlPattern)

	// Find the first match in the text
	match := re.FindString(text)

	if match == "" {
		return "", fmt.Errorf("no url found in provided text")
	}
	return match, nil
}

// 获取重定向地址
func HttpGetRedirectURL(url string) (string, error) {
	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 获取重定向URL
	redirectURL := resp.Request.URL
	if redirectURL.String() == "" {
		return "", fmt.Errorf("empty redirect URL")
	}
	return redirectURL.String(), nil
}

func DownloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error download %s: %v", filepath, err)
	}
	defer resp.Body.Close()

	f, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("error create file %s: %v", filepath, err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("error writ to file %s: %v", filepath, err)
	}
	return nil
}
